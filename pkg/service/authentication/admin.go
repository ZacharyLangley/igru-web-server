package authentication

import (
	"crypto/rand"
	"encoding/base64"
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
	adminEmail = "admin"
	adminName  = "admin"
)

func randomPassword() (string, error) {
	randBuffer := make([]byte, 16)
	if _, err := rand.Read(randBuffer); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(randBuffer), nil
}

func CreateAdmin(ctx context.Context, conn *database.Pool) error {
	var password string
	err := conn.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(conn)
		// Check if account already exists
		_, err := queries.GetUserByEmail(ctx, adminEmail)
		found := errors.Is(err, pgx.ErrNoRows)
		if err != nil && !found {
			return err
		}
		if found {
			return nil
		}
		// Create new admin password
		password, err = randomPassword()
		if err != nil {
			return err
		}
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
	ctx.L().Info("created new admin account", zap.String("email", adminEmail), zap.String("password", password))
	return nil
}
