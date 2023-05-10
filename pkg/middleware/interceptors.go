package middleware

import (
	gocontext "context"
	"errors"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/bufbuild/connect-go"
	otelconnect "github.com/bufbuild/connect-opentelemetry-go"
	"go.uber.org/zap"
)

func Interceptors(serviceName string) []connect.Interceptor {
	return []connect.Interceptor{
		otelconnect.NewInterceptor(),
		logRequest(),
	}
}

func logRequest() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(baseCtx gocontext.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			ctx := context.New(baseCtx).Named("service").WithFields(
				zap.String("procedure", req.Spec().Procedure),
				zap.String("peer", req.Peer().Addr),
			).WithTrace()
			res, err := next(ctx, req)
			if err != nil {
				cErr := &connect.Error{}
				ok := errors.As(err, &cErr)
				if ok {
					ctx.L().Error("fail processing request", zap.Error(cErr))
				}
			}
			return res, err
		})
	}
}
