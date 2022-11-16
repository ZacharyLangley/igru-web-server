package authentication

import (
	"database/sql"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/middleware"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/bufbuild/connect-go"
)

func New(conn *sql.DB) *Service {
	return &Service{
		conn: conn,
	}
}

type Service struct {
	conn *sql.DB
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
