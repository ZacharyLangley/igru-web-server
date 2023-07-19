package authentication

import (
	"errors"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
)

const (
	adminEmail = "admin@admin"
	adminName  = "admin"

	//nolint: gosec
	defaultAdminPassword = "GankyCumdumpster69"
)

func CreateAdmin(ctx context.Context, conn *database.Pool) error {
	password := defaultAdminPassword
	err := conn.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(conn)
		// Check if account already exists
		_, err := queries.GetUserByEmail(ctx, adminEmail)
		found := !errors.Is(err, pgx.ErrNoRows)
		if err != nil && found {
			return err
		}
		if found {
			return nil
		}
		// Create new admin password
		hash, salt, err := generateCred(password)
		if err != nil {
			return err
		}
		// Create new admin group
		group, err := queries.CreateUserGroup(ctx, adminEmail)
		if err != nil {
			return err
		}
		// Create new admin account
		params := models.CreateUserParams{
			GroupID:    group.ID,
			Email:      adminEmail,
			FullName:   pgtype.Text{Valid: true, String: adminName},
			GlobalRole: pgtype.Int4{Valid: true, Int32: int32(v1.GroupRole_GROUP_ROLE_ADMIN)},
			Salt:       salt,
			Hash:       hash,
		}
		user, err := queries.CreateUser(ctx, params)
		if err != nil {
			return err
		}
		// Associate admin user with admin group
		return queries.AddGroupMember(ctx, models.AddGroupMemberParams{
			UserID:  user.ID,
			GroupID: group.ID,
			Role:    int32(v1.GroupRole_GROUP_ROLE_ADMIN),
		})
	})
	if err != nil {
		return err
	}
	if password != "" {
		ctx.L().Info("created new admin account", zap.String("email", adminEmail), zap.String("password", password))
	} else {
		ctx.L().Info("admin account already created")
	}
	return nil
}
