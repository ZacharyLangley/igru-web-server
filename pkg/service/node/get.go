package node

import (
	gocontext "context"
	"errors"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/node"
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/node/v1"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) GetNodes(baseCtx gocontext.Context, req *connect.Request[v1.GetNodesRequest]) (*connect.Response[v1.GetNodesResponse], error) {
	var err error
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetNodesResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	groupID := uuid.NullUUID{}
	if req.Msg.GroupId != nil {
		groupID.UUID, err = uuid.Parse(*req.Msg.GroupId)
		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid groupID"))
		} else {
			groupID.Valid = true
		}
	}
	var allowed bool
	if groupID.Valid {
		// Check group permissions
		allowed, err = s.checker.AssertRead(ctx, req, *req.Msg.GroupId)
		if err != nil {
			return nil, err
		}
	} else {
		// Check global permissions
		allowed, err = s.checker.AssertGlobalRead(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	if !allowed {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("access denied"))
	}
	var nodes []models.Node
	sensors := make(map[string][]models.Sensor)
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		queries := models.New(tx)
		params := models.GetNodesParams{
			OwnedBy: groupID,
			Limit:   req.Msg.Pagination.Length,
			Offset:  req.Msg.Pagination.Cursor,
		}
		nodes, err = queries.GetNodes(ctx, params)
		if err != nil {
			return err
		}
		for _, node := range nodes {
			sensors[node.MacAddress], err = queries.GetNodeSensors(ctx, node.MacAddress)
			if err != nil {
				return err
			}
		}
		return err
	}); err != nil {
		return nil, err
	}
	res.Msg.Nodes = make([]*v1.Node, 0, len(nodes))
	for _, node := range nodes {
		newNode := v1.Node{
			MacAddress: node.MacAddress,
			Name:       node.Name,
			CreatedAt:  timestamppb.New(node.CreatedAt),
		}

		if node.OwnedBy.Valid {
			newNode.OwnedBy = node.OwnedBy.UUID.String()
		}
		if err := node.CustomLabels.AssignTo(newNode.CustomLabels); err != nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
		if node.UpdatedAt.Valid {
			newNode.UpdatedAt = timestamppb.New(node.UpdatedAt.Time)
		}
		if node.AdoptedAt.Valid {
			newNode.AdoptedAt = timestamppb.New(node.AdoptedAt.Time)
		}
		res.Msg.Nodes = append(res.Msg.Nodes, &newNode)
	}
	return res, nil
}

func (s *Service) GetNode(gocontext.Context, *connect.Request[v1.GetNodeRequest]) (*connect.Response[v1.GetNodeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("node.v1.NodeService.GetNode is not implemented"))
}
