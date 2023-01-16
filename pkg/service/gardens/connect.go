package gardens

import (
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/auth"
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
	pool    *database.Pool
	checker *auth.Checker
	gardensv1connect.UnimplementedGardensServiceHandler
	gardensv1connect.UnimplementedPlantsServiceHandler
}

func (s *Service) Register(mux *http.ServeMux) {
	interceptors := middleware.Interceptors("garden-service")
	mux.Handle(gardensv1connect.NewGardensServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
	mux.Handle(gardensv1connect.NewPlantsServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
	mux.Handle(gardensv1connect.NewStrainsServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
	mux.Handle(gardensv1connect.NewRecipesServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
}
