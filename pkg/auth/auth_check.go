package auth

import (
	"fmt"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	authenticationv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
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
