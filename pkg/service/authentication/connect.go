package authentication

import (
	"net/http"

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
	pool *database.Pool
	authenticationv1connect.UnimplementedUserServiceHandler
	authenticationv1connect.UnimplementedGroupServiceHandler
}

func (s *Service) Register(mux *http.ServeMux) {
	interceptors := middleware.Interceptors("authorization-service")
	mux.Handle(authenticationv1connect.NewUserServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
	mux.Handle(authenticationv1connect.NewGroupServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
}
