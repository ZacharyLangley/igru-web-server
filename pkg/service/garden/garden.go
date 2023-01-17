package garden

import (
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/auth"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	"github.com/ZacharyLangley/igru-web-server/pkg/middleware"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/garden/v1/gardenv1connect"
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
	gardenv1connect.UnimplementedGardenServiceHandler
	gardenv1connect.UnimplementedPlantServiceHandler
	gardenv1connect.UnimplementedStrainServiceHandler
	gardenv1connect.UnimplementedRecipeServiceHandler
}

func (s *Service) Register(mux *http.ServeMux) {
	interceptors := middleware.Interceptors("garden-service")
	mux.Handle(gardenv1connect.NewGardenServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
	mux.Handle(gardenv1connect.NewPlantServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
	mux.Handle(gardenv1connect.NewStrainServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
	mux.Handle(gardenv1connect.NewRecipeServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
}
