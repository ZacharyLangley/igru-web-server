package authentication

import (
	"database/sql"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
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
	mux.Handle(authenticationv1connect.NewUserServiceHandler(s))
}
