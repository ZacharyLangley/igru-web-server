package authentication

import (
	"net/http"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	"github.com/ZacharyLangley/igru-web-server/pkg/middleware"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/bufbuild/connect-go"
)

func New(pool *database.Pool) *Service {
	return &Service{
		pool: pool,
	}
}

type Service struct {
	SessionDuration time.Duration
	pool            *database.Pool
	authenticationv1connect.UnimplementedUserServiceHandler
	authenticationv1connect.UnimplementedGroupServiceHandler
	authenticationv1connect.UnimplementedSessionServiceHandler
}

func (s *Service) Register(mux *http.ServeMux) {
	interceptors := middleware.Interceptors("authorization-service")
	mux.Handle(authenticationv1connect.NewUserServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
	mux.Handle(authenticationv1connect.NewGroupServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
	mux.Handle(authenticationv1connect.NewSessionServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
}
