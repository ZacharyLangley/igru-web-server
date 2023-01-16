package auth

import (
	"fmt"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	authenticationv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

func GetGroups(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (map[uuid.UUID]struct{}, error) {
	// Get user groups
	query := models.New(tx)
	userGroups, err := query.GetUserGroups(ctx, models.GetUserGroupsParams{
		UserID: userID,
		Limit:  100,
	})
	if err != nil {
		return nil, fmt.Errorf("failed getting user groups: %w", err)
	}
	groupIDs := make(map[uuid.UUID]struct{})
	for _, group := range userGroups {
		groupIDs[group.ID] = struct{}{}
	}
	return groupIDs, nil
}

type Checker struct {
	authenticationv1connect.SessionServiceClient
}

func (c Checker) AssertAll(ctx context.Context, req interface {
	Header() http.Header
}, groupID string, roles ...authenticationv1.GroupRole) (bool, error) {
	request := connect.NewRequest(&authenticationv1.CheckSessionPermissionsRequest{})
	if len(roles) == 0 {
		return false, nil
	}
	request.Msg.Requests = make([]*authenticationv1.PermissionRequest, len(roles))
	for i, role := range roles {
		request.Msg.Requests[i] = &authenticationv1.PermissionRequest{
			GroupId: groupID,
			Role:    role,
		}
	}
	res, err := c.SessionServiceClient.CheckSessionPermissions(ctx, request)
	if err != nil {
		return false, err
	}
	for _, response := range res.Msg.Responses {
		if !response.Allowed {
			return false, nil
		}
	}
	return true, nil
}
func (c Checker) AssertAny(ctx context.Context, req interface {
	Header() http.Header
}, groupID string, roles ...authenticationv1.GroupRole) (bool, error) {
	request := connect.NewRequest(&authenticationv1.CheckSessionPermissionsRequest{})
	if len(roles) == 0 {
		return false, nil
	}
	request.Msg.Requests = make([]*authenticationv1.PermissionRequest, len(roles))
	for i, role := range roles {
		request.Msg.Requests[i] = &authenticationv1.PermissionRequest{
			GroupId: groupID,
			Role:    role,
		}
	}
	res, err := c.SessionServiceClient.CheckSessionPermissions(ctx, request)
	if err != nil {
		return false, err
	}
	for _, response := range res.Msg.Responses {
		if !response.Allowed {
			return true, nil
		}
	}
	return false, nil
}

func Check(ctx context.Context, tx pgx.Tx, userID uuid.UUID, requests []*authenticationv1.PermissionRequest) ([]*authenticationv1.PermissionResponse, error) {
	// Get user groups
	query := models.New(tx)
	userGroups, err := query.GetUserGroups(ctx, models.GetUserGroupsParams{
		UserID: userID,
		Limit:  100,
	})
	if err != nil {
		return nil, fmt.Errorf("failed getting user groups: %w", err)
	}
	groupIDs := make(map[uuid.UUID]struct{})
	for _, group := range userGroups {
		groupIDs[group.ID] = struct{}{}
	}
	// Prepare output
	output := make([]*authenticationv1.PermissionResponse, len(requests))
	for i := range output {
		output[i] = &authenticationv1.PermissionResponse{
			Request: requests[i],
		}
	}
	// Check if requests belong to user group
	for i, req := range requests {
		groupID, err := uuid.Parse(req.GroupId)
		if err != nil {
			return nil, fmt.Errorf("error parsing group UUID on request %d: %w", i, err)
		}
		_, output[i].Allowed = groupIDs[groupID]
	}
	return output, nil
}
