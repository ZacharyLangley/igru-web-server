package node

import (
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/auth"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	"github.com/ZacharyLangley/igru-web-server/pkg/middleware"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/node/v1/nodev1connect"
	"github.com/bufbuild/connect-go"
)

func New(pool *database.Pool, checker *auth.Checker) *Service {
	return &Service{
		pool:    pool,
		checker: checker,
	}
}

type Service struct {
	pool    *database.Pool
	checker *auth.Checker
	nodev1connect.UnimplementedNodeServiceHandler
}

func (s *Service) Register(mux *http.ServeMux) {
	interceptors := middleware.Interceptors("node-service")
	mux.Handle(nodev1connect.NewNodeServiceHandler(s,
		connect.WithInterceptors(interceptors...)))
}
