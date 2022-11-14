package authentication

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/go-redis/redis/v8"
)

func New(conn *sql.DB, cache *redis.Client) *Service {
	return &Service{
		conn:            conn,
		cache:           cache,
		sessionDuration: time.Minute * 10,
	}
}

type Service struct {
	conn            *sql.DB
	cache           *redis.Client
	sessionDuration time.Duration
	authenticationv1connect.UnimplementedUserServiceHandler
}

func (s *Service) Register(mux *http.ServeMux) {
	mux.Handle(authenticationv1connect.NewUserServiceHandler(s))
}
