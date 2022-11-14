package authentication

import (
	gocontext "context"
	"database/sql"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/bufbuild/connect-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.uber.org/zap"
)

func New(conn *sql.DB) *Service {
	return &Service{
		conn: conn,
	}
}

type Service struct {
	conn *sql.DB
	authenticationv1connect.UnimplementedUserServiceHandler
}

func (s *Service) Register(mux *http.ServeMux) {
	mux.Handle(authenticationv1connect.NewUserServiceHandler(s,
		connect.WithInterceptors(
			connect.UnaryInterceptorFunc(logRequest),
			connect.UnaryInterceptorFunc(otelInterceptor),
		)))
}

const serviceName = "authorization-service"

func otelInterceptor(next connect.UnaryFunc) connect.UnaryFunc {
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

func logRequest(next connect.UnaryFunc) connect.UnaryFunc {
	return connect.UnaryFunc(func(baseCtx gocontext.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		ctx := context.New(baseCtx).Named("service").WithFields(
			zap.String("procedure", req.Spec().Procedure),
			zap.String("peer", req.Peer().Addr),
		)
		ctx.L().Info("start")
		res, err := next(ctx, req)
		ctx.L().Info("end")
		return res, err
	})
}
