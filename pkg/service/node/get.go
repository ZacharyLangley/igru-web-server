package node

import (
	gocontext "context"
	"errors"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/node/v1"
	"github.com/bufbuild/connect-go"
)

func (s *Service) GetNodes(baseCtx gocontext.Context, req *connect.Request[v1.GetNodesRequest]) (*connect.Response[v1.GetNodesResponse], error) {
	var err error
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetNodesResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	allowed, err := s.checker.AssertRead(ctx, req, "TODO: add global admin group")
	if err != nil {
		return nil, err
	}
	if !allowed {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("access denied"))
	}
	return res, nil
}

func (s *Service) GetNode(gocontext.Context, *connect.Request[v1.GetNodeRequest]) (*connect.Response[v1.GetNodeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("node.v1.NodeService.GetNode is not implemented"))
}
