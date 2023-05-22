package otlplogs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel/sdk/resource"
	collpb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	v1 "go.opentelemetry.io/proto/otlp/common/v1"
	lpb "go.opentelemetry.io/proto/otlp/logs/v1"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

const (
	TraceIDFieldName = "trace_id"
	SpanIDFieldName  = "span_id"
)

type Core struct {
	Level            zapcore.LevelEnabler
	resource         *resource.Resource
	additionalFields []zapcore.Field
	syncChan         chan struct{}
	inputBuffer      chan *lpb.LogRecord
}

// NewNopCore returns a no-op Core.
func NewOTELCore(ctx context.Context, conn *grpc.ClientConn, opts ...ProviderOption) Core {
	batch := &batcher[*lpb.LogRecord]{
		1000,
		time.Second * 5,
	}
	syncChan := make(chan struct{})
	inputBuffer := make(chan *lpb.LogRecord)
	client := collpb.NewLogsServiceClient(conn)
	core := Core{
		syncChan:    syncChan,
		inputBuffer: inputBuffer,
	}
	for _, opt := range opts {
		core = opt.apply(core)
	}
	go core.start(ctx, client, batch.Start(ctx, syncChan, inputBuffer))
	return core
}
func (c Core) start(ctx context.Context, client collpb.LogsServiceClient, buffer <-chan []*lpb.LogRecord) {
	for {
		select {
		case <-ctx.Done():
			return
		case logs, ok := <-buffer:
			if !ok {
				return
			}
			resp, err := client.Export(context.Background(), &collpb.ExportLogsServiceRequest{
				ResourceLogs: []*lpb.ResourceLogs{
					{
						ScopeLogs: []*lpb.ScopeLogs{
							{
								LogRecords: logs,
							},
						},
					},
				},
			})
			if err != nil {
				log.Printf("Failed to send logs: %#v %#v", resp, err)
			}
		}
	}
}
func (c Core) Enabled(level zapcore.Level) bool {
	return c.Level.Enabled(level)
}
func (c Core) With(fields []zapcore.Field) zapcore.Core {
	clone := c.clone()
	clone.additionalFields = append(clone.additionalFields, fields...)
	return clone
}

func (c Core) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Level.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c Core) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	record := lpb.LogRecord{
		TimeUnixNano:         uint64(time.Now().UnixNano()),
		ObservedTimeUnixNano: uint64(ent.Time.UnixNano()),
		SeverityText:         ent.Level.String(),
		Body: &v1.AnyValue{
			Value: &v1.AnyValue_StringValue{
				StringValue: ent.Message,
			},
		},
	}
	for _, field := range append(c.additionalFields, fields...) {
		switch field.Key {
		case SpanIDFieldName:
			//nolint: forcetypeassert
			record.SpanId = field.Interface.([]byte)
		case TraceIDFieldName:
			//nolint: forcetypeassert
			record.TraceId = field.Interface.([]byte)
		default:
			var value v1.AnyValue
			switch field.Type {
			case zapcore.BoolType:
				//nolint: forcetypeassert
				value.Value = &v1.AnyValue_BoolValue{
					BoolValue: field.Interface.(bool),
				}
			case zapcore.Float32Type:
				//nolint: forcetypeassert
				value.Value = &v1.AnyValue_DoubleValue{
					DoubleValue: float64(field.Interface.(float32)),
				}
			case zapcore.Float64Type:
				//nolint: forcetypeassert
				value.Value = &v1.AnyValue_DoubleValue{
					DoubleValue: field.Interface.(float64),
				}
			case zapcore.Int8Type:
				fallthrough
			case zapcore.Int16Type:
				fallthrough
			case zapcore.Int32Type:
				fallthrough
			case zapcore.Int64Type:
				fallthrough
			case zapcore.Uint8Type:
				fallthrough
			case zapcore.Uint16Type:
				fallthrough
			case zapcore.Uint32Type:
				fallthrough
			case zapcore.Uint64Type:
				value.Value = &v1.AnyValue_IntValue{
					IntValue: field.Integer,
				}
			case zapcore.ErrorType:
				value.Value = &v1.AnyValue_StringValue{
					StringValue: field.String,
				}
			default:
				value.Value = &v1.AnyValue_StringValue{
					StringValue: field.String,
				}
			}
			record.Attributes = append(record.Attributes, &v1.KeyValue{
				Key:   field.Key,
				Value: &value,
			})
		}
	}
	switch ent.Level {
	case zapcore.PanicLevel:
		record.SeverityNumber = lpb.SeverityNumber_SEVERITY_NUMBER_FATAL
	case zapcore.DPanicLevel:
		record.SeverityNumber = lpb.SeverityNumber_SEVERITY_NUMBER_FATAL
	case zapcore.DebugLevel:
		record.SeverityNumber = lpb.SeverityNumber_SEVERITY_NUMBER_DEBUG
	case zapcore.ErrorLevel:
		record.SeverityNumber = lpb.SeverityNumber_SEVERITY_NUMBER_ERROR
	case zapcore.FatalLevel:
		record.SeverityNumber = lpb.SeverityNumber_SEVERITY_NUMBER_FATAL
	case zapcore.InfoLevel:
		record.SeverityNumber = lpb.SeverityNumber_SEVERITY_NUMBER_INFO
	case zapcore.WarnLevel:
		record.SeverityNumber = lpb.SeverityNumber_SEVERITY_NUMBER_WARN
	default:
		return invalidSecurityLevelError{ent.Level}
	}
	c.inputBuffer <- &record
	return nil
}

type invalidSecurityLevelError struct {
	Level zapcore.Level
}

func (i invalidSecurityLevelError) Error() string {
	return fmt.Sprintf("invalid severity level: %s", i.Level)
}

func (c Core) Sync() error {
	c.syncChan <- struct{}{}
	return nil
}

func (c *Core) clone() *Core {
	return &Core{
		Level:            c.Level,
		resource:         c.resource,
		additionalFields: c.additionalFields,
		syncChan:         c.syncChan,
		inputBuffer:      c.inputBuffer,
	}
}
