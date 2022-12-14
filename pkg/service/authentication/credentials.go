package authentication

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
	"time"

	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const sessionDuration = time.Minute * 5

var errUnauthorizedUser = errors.New("could not authenticate user")

func generateCred(password string) (string, string, error) {
	rawSalt := make([]byte, 128)
	if _, err := rand.Read(rawSalt); err != nil {
		return "", "", err
	}
	salt := base64.StdEncoding.EncodeToString(rawSalt)
	digest := append(rawSalt, []byte(password)...)
	rawHash, err := bcrypt.GenerateFromPassword(digest, 0)
	if err != nil {
		return "", "", err
	}
	hash := base64.StdEncoding.EncodeToString(rawHash)
	return hash, salt, nil
}

func checkHash(user models.User, password string) bool {
	salt, err := base64.StdEncoding.DecodeString(user.Salt)
	if err != nil {
		return false
	}
	userDigest := append(salt, []byte(password)...)
	hashDigest, err := base64.StdEncoding.DecodeString(user.Hash)
	if err != nil {
		return false
	}
	if err := bcrypt.CompareHashAndPassword(hashDigest, userDigest); err != nil {
		// TODO check for non-mismatched errors
		return false
	}
	return true
}

func AddSessionToken(h http.Header, token string) {
	h.Add("Authentication", "Bearer: "+token)
}

func ExtractSessionToken(h http.Header) (uuid.UUID, error) {
	authHeader := h.Get("Authentication")
	if authHeader == "" {
		return uuid.Nil, errUnauthorizedUser
	}
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 {
		return uuid.Nil, errUnauthorizedUser
	}
	return uuid.Parse(authParts[1])
}
