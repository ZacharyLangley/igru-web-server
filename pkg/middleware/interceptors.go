package middleware

import (
	gocontext "context"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/bufbuild/connect-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.uber.org/zap"
)

func Interceptors(serviceName string) []connect.Interceptor {
	return []connect.Interceptor{
		connect.UnaryInterceptorFunc(logRequest()),
		connect.UnaryInterceptorFunc(otelInterceptor(serviceName)),
	}
}

func otelInterceptor(serviceName string) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(baseCtx gocontext.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			ctx, span := otel.Tracer(serviceName).Start(baseCtx, req.Spec().Procedure)
			res, err := next(ctx, req)
			if err != nil {
				if cErr, ok := err.(*connect.Error); ok {
					span.SetStatus(codes.Error, cErr.Message())
				}
			}
			span.End()
			return res, err
		})
	}
}

func logRequest() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(baseCtx gocontext.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			ctx := context.New(baseCtx).Named("service").WithFields(
				zap.String("procedure", req.Spec().Procedure),
				zap.String("peer", req.Peer().Addr),
			)
			ctx.L().Debug("start")
			res, err := next(ctx, req)
			if err != nil {
				if cErr, ok := err.(*connect.Error); ok {
					ctx.L().Error("fail processing request", zap.Error(cErr))
				}
			}
			ctx.L().Debug("end")
			return res, err
		})
	}
}
