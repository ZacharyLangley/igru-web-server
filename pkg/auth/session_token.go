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

func DecodeToken(token string) (*models.Session, error) {
	jsonMessage, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	var message sessionFormat
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		return nil, err
	}
	var output models.Session
	output.ID = message.ID
	output.UserID = message.UserID
	output.CreatedAt = message.IssuedAt
	output.ExpiredAt = message.ExpiredAt
	return &output, nil
}

func EncodeToken(sess models.Session) (string, error) {
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
