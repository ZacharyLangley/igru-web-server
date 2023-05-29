package node

import (
	gocontext "context"
	"encoding/json"
	"errors"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/node"
	authenticationv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/node/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/common"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) GetNodes(baseCtx gocontext.Context, req *connect.Request[v1.GetNodesRequest]) (*connect.Response[v1.GetNodesResponse], error) {
	var err error
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetNodesResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	groupID, err := s.checker.AssertAny(ctx,
		req,
		req.Msg.GroupId,
		authenticationv1.GroupRole_GROUP_ROLE_ADMIN,
		authenticationv1.GroupRole_GROUP_ROLE_READ_ONLY,
	)
	if err != nil {
		return nil, err
	}
	var nodes []models.Node
	sensors := make(map[string][]models.Sensor)
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		queries := models.New(tx)
		params := models.GetNodesParams{
			OwnedBy: database.NewFromNullableUUID(groupID),
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
			CreatedAt:  timestamppb.New(node.CreatedAt.Time),
		}
		if node.OwnedBy.Valid {
			newNode.OwnedBy = uuid.UUID(node.OwnedBy.Bytes).String()
		}
		node.CustomLabels, err = json.Marshal(newNode.CustomLabels)
		if err != nil {
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
	// TODO: Fill in GetNode
	//nolint: all
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("node.v1.NodeService.GetNode is not implemented"))
}
