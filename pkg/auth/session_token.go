package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type sessionFormat struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"uid"`
	IssuedAt  time.Time `json:"iss"`
	ExpiredAt time.Time `json:"exp"`
}

func DecodeToken(token string) (*models.Session, error) {
	jsonMessage, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, fmt.Errorf("failed base64 decode: %w", err)
	}
	var message sessionFormat
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		return nil, fmt.Errorf("failed json decode: %w", err)
	}
	var output models.Session
	output.ID = pgtype.UUID{
		Bytes: message.ID,
		Valid: true,
	}
	output.UserID = pgtype.UUID{
		Bytes: message.UserID,
		Valid: true,
	}
	output.CreatedAt = pgtype.Timestamp{
		Time:  message.IssuedAt,
		Valid: true,
	}
	output.ExpiredAt = pgtype.Timestamp{
		Time:  message.ExpiredAt,
		Valid: true,
	}
	return &output, nil
}

func EncodeToken(sess models.Session) (string, error) {
	jsonMessage, err := json.Marshal(sessionFormat{
		ID:        sess.ID.Bytes,
		UserID:    sess.UserID.Bytes,
		IssuedAt:  sess.CreatedAt.Time,
		ExpiredAt: sess.ExpiredAt.Time,
	})
	if err != nil {
		return "", fmt.Errorf("failed json encoding: %w", err)
	}
	return base64.URLEncoding.EncodeToString(jsonMessage), nil
}
