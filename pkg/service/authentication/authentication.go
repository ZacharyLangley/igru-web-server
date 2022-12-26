package authentication

import (
	"net/http"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/auth"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	"github.com/ZacharyLangley/igru-web-server/pkg/middleware"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/bufbuild/connect-go"
)

func New(pool *database.Pool, authServer *auth.Server) *Service {
	return &Service{
		pool:       pool,
		authServer: authServer,
	}
}

type authService interface {
	CreateSessionJWT(models.Session) (string, error)
}

type Service struct {
	SessionDuration time.Duration
	pool            *database.Pool
	authServer      *auth.Server
	authenticationv1connect.UnimplementedUserServiceHandler
	authenticationv1connect.UnimplementedGroupServiceHandler
	authenticationv1connect.UnimplementedAuthServiceHandler
}

func (s *Service) Register(mux *http.ServeMux) {
	interceptors := middleware.Interceptors("authorization-service")
	mux.Handle(authenticationv1connect.NewUserServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
	mux.Handle(authenticationv1connect.NewGroupServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
	mux.Handle(authenticationv1connect.NewAuthServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
	mux.Handle("/oauth2/v3/certs", s.authServer)
}
