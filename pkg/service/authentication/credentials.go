package authentication

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"

	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	"golang.org/x/crypto/argon2"
)

var errUnauthorizedUser = errors.New("could not authenticate user")

func generateCred(password string) (string, string, error) {
	rawSalt := make([]byte, 128)
	if _, err := rand.Read(rawSalt); err != nil {
		return "", "", err
	}
	salt := base64.StdEncoding.EncodeToString(rawSalt)
	rawHash := argon2.IDKey([]byte(password), rawSalt, 1, 64*1024, 4, 32)
	hash := base64.StdEncoding.EncodeToString(rawHash)
	return hash, salt, nil
}

func checkHash(user models.User, password string) bool {
	rawSalt, err := base64.StdEncoding.DecodeString(user.Salt)
	if err != nil {
		return false
	}
	hashDigest, err := base64.StdEncoding.DecodeString(user.Hash)
	if err != nil {
		return false
	}
	rawHash := argon2.IDKey([]byte(password), rawSalt, 1, 64*1024, 4, 32)
	return bytes.Equal(hashDigest, rawHash)
}

func AddSessionToken(h http.Header, token string) {
	h.Add("SESSION", token)
}

func ExtractSessionToken(h http.Header) (string, error) {
	authHeader := h.Get("SESSION")
	if authHeader == "" {
		return "", errUnauthorizedUser
	}
	return authHeader, nil
}
