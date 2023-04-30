package node

import (
	gocontext "context"
	"errors"

	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/node/v1"
	connect_go "github.com/bufbuild/connect-go"
)

func (s *Service) UpdateNode(gocontext.Context, *connect_go.Request[v1.UpdateNodeRequest]) (*connect_go.Response[v1.UpdateNodeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("node.v1.NodeService.UpdateNode is not implemented"))
}
