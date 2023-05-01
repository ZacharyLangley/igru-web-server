package auth

import (
	"errors"
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

func (c Checker) AssertAny(ctx context.Context, req interface {
	Header() http.Header
}, groupID *string, roles ...authenticationv1.GroupRole) (uuid.NullUUID, error) {
	// Parse Group ID
	var err error
	var output uuid.NullUUID
	if groupID != nil {
		output.UUID, err = uuid.Parse(*groupID)
		if err != nil {
			return output, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid groupID"))
		} else {
			output.Valid = true
		}
	}
	// Prep group permissions check
	request := connect.NewRequest(&authenticationv1.CheckSessionPermissionsRequest{})
	request.Msg.Requests = make([]*authenticationv1.PermissionRequest, 0, len(roles)*2)
	for _, role := range roles {
		request.Msg.Requests = append(request.Msg.Requests,
			&authenticationv1.PermissionRequest{
				Role: role,
			},
		)
	}
	if output.Valid {
		for _, role := range roles {
			request.Msg.Requests = append(request.Msg.Requests,
				&authenticationv1.PermissionRequest{
					Role: role,
				},
			)
		}
	}
	// Send request
	res, err := c.SessionServiceClient.CheckSessionPermissions(ctx, request)
	if err != nil {
		return uuid.NullUUID{}, connect.NewError(connect.CodeInternal, fmt.Errorf("failed checking permissions: %w", err))
	}
	// Check response
	for _, response := range res.Msg.Responses {
		if response.Allowed {
			return output, nil
		}
	}
	return uuid.NullUUID{}, connect.NewError(connect.CodePermissionDenied, errors.New("access denied"))
}

func (c Checker) AssertRead(ctx context.Context, req interface {
	Header() http.Header
}, groupID uuid.NullUUID) (bool, error) {
	request := connect.NewRequest(&authenticationv1.CheckSessionPermissionsRequest{})
	request.Msg.Requests = []*authenticationv1.PermissionRequest{
		{
			Role: authenticationv1.GroupRole_GROUP_ROLE_ADMIN,
		},
		{
			Role: authenticationv1.GroupRole_GROUP_ROLE_READ_ONLY,
		},
	}
	if groupID.Valid {
		validGroupID := groupID.UUID.String()
		request.Msg.Requests = append(request.Msg.Requests,
			&authenticationv1.PermissionRequest{
				GroupId: &validGroupID,
				Role:    authenticationv1.GroupRole_GROUP_ROLE_ADMIN,
			},
			&authenticationv1.PermissionRequest{
				GroupId: &validGroupID,
				Role:    authenticationv1.GroupRole_GROUP_ROLE_READ_ONLY,
			},
		)
	}
	res, err := c.SessionServiceClient.CheckSessionPermissions(ctx, request)
	if err != nil {
		return false, fmt.Errorf("failed checking permissions: %w", err)
	}
	for _, response := range res.Msg.Responses {
		if response.Allowed {
			return true, nil
		}
	}
	return false, nil
}

func (c Checker) AssertWrite(ctx context.Context, req interface {
	Header() http.Header
}, groupID uuid.NullUUID) (bool, error) {
	request := connect.NewRequest(&authenticationv1.CheckSessionPermissionsRequest{})
	request.Msg.Requests = []*authenticationv1.PermissionRequest{
		{
			Role: authenticationv1.GroupRole_GROUP_ROLE_ADMIN,
		},
	}
	if groupID.Valid {
		validGroupID := groupID.UUID.String()
		request.Msg.Requests = append(request.Msg.Requests,
			&authenticationv1.PermissionRequest{
				GroupId: &validGroupID,
				Role:    authenticationv1.GroupRole_GROUP_ROLE_ADMIN,
			},
		)
	}
	res, err := c.SessionServiceClient.CheckSessionPermissions(ctx, request)
	if err != nil {
		return false, fmt.Errorf("failed checking permissions: %w", err)
	}
	for _, response := range res.Msg.Responses {
		if response.Allowed {
			return true, nil
		}
	}
	return false, nil
}
