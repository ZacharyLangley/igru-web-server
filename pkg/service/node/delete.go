package node

import (
	gocontext "context"
	"errors"

	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/node/v1"
	connect_go "github.com/bufbuild/connect-go"
)

func (s *Service) DeleteNode(gocontext.Context, *connect_go.Request[v1.DeleteNodeRequest]) (*connect_go.Response[v1.DeleteNodeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("node.v1.NodeService.DeleteNode is not implemented"))
}
