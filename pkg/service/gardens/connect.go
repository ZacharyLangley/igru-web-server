package gardens

import (
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	"github.com/ZacharyLangley/igru-web-server/pkg/middleware"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/gardens/v1/gardensv1connect"
	"github.com/bufbuild/connect-go"
)

func New(pool *database.Pool) *Service {
	return &Service{
		pool: pool,
	}
}

type Service struct {
	pool *database.Pool
	gardensv1connect.UnimplementedGardensServiceHandler
}

func (s *Service) Register(mux *http.ServeMux) {

	interceptors := middleware.Interceptors("authorization-service")
	mux.Handle(gardensv1connect.NewGardensServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
}
