package auth

import (
	"encoding/base64"
	"encoding/json"

	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
)

func VerifySessionJWT(token string) (*models.Session, error) {
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
