package context

import (
	"context"

	"github.com/ZacharyLangley/igru-web-server/pkg/config/otlplogs"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Context interface {
	context.Context
	L() *zap.Logger
	Named(name string) Context
	WithFields(fields ...zap.Field) Context
	WithTrace() Context
	AddEvent(string)
}

type internalContext struct {
	//nolint:containedctx
	context.Context
	logger *zap.Logger
}

func (c *internalContext) Named(name string) Context {
	return &internalContext{
		Context: c.Context,
		logger:  c.L().Named(name),
	}
}

func (c *internalContext) WithFields(fields ...zap.Field) Context {
	return &internalContext{
		Context: c.Context,
		logger:  c.L().With(fields...),
	}
}

func (c *internalContext) WithTrace() Context {
	if spanCtx := trace.SpanContextFromContext(c); spanCtx.IsValid() {
		tID := spanCtx.TraceID()
		sID := spanCtx.SpanID()
		return c.WithFields(
			zap.Binary(otlplogs.SpanIDFieldName, sID[:]),
			zap.Binary(otlplogs.TraceIDFieldName, tID[:]),
		)
	}
	zap.L().Warn("Could not find trace for logger")
	return c
}

func (c *internalContext) AddEvent(msg string) {
	if span := trace.SpanFromContext(c); span != nil {
		span.AddEvent(msg)
	}
}

func (c *internalContext) L() *zap.Logger {
	if c.logger != nil {
		return c.logger
	}
	return zap.L()
}

func New(ctx context.Context) Context {
	if nested, ok := ctx.(Context); ok {
		return nested
	}
	return &internalContext{
		Context: ctx,
	}
}

func WithLogger(ctx context.Context, logger *zap.Logger) Context {
	if nested, ok := ctx.(*internalContext); ok {
		return &internalContext{
			Context: nested.Context,
			logger:  logger,
		}
	}
	return &internalContext{
		Context: ctx,
		logger:  logger,
	}
}
