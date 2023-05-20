package node

import (
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	"github.com/ZacharyLangley/igru-web-server/pkg/middleware"
	authenticationv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/node/v1/nodev1connect"
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
	nodev1connect.UnimplementedNodeServiceHandler
}

func (s *Service) Register(mux *http.ServeMux) {
	interceptors := middleware.Interceptors("node-service")
	mux.Handle(nodev1connect.NewNodeServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
}
