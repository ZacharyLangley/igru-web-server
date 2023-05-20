package garden

import (
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	"github.com/ZacharyLangley/igru-web-server/pkg/middleware"
	authenticationv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/garden/v1/gardenv1connect"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
)

func New(pool *database.Pool, checker authChecker) *Service {
	return &Service{
		pool:    pool,
		checker: checker,
	}
}

type authChecker interface {
	AssertAny(ctx context.Context, req interface {
		Header() http.Header
	}, groupID *string, roles ...authenticationv1.GroupRole) (uuid.NullUUID, error)
}

type Service struct {
	pool    *database.Pool
	checker authChecker
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
