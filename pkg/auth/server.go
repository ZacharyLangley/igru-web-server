package auth

import (
	"encoding/base64"
	"encoding/json"
	"time"

	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	"github.com/google/uuid"
)

type sessionFormat struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"uid"`
	IssuedAt  time.Time `json:"iss"`
	ExpiredAt time.Time `json:"exp"`
}

func CreateSession(sess models.Session) (string, error) {
	jsonMessage, err := json.Marshal(sessionFormat{
		ID:        sess.ID,
		UserID:    sess.UserID,
		IssuedAt:  sess.CreatedAt,
		ExpiredAt: sess.ExpiredAt,
	})
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(jsonMessage), nil
}
