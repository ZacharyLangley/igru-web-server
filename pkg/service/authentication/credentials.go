package authentication

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
	"time"

	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
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
	h.Add("SESSION", token)
}

func ExtractSessionToken(h http.Header) (string, error) {
	authHeader := h.Get("SESSION")
	if authHeader == "" {
		return "", errUnauthorizedUser
	}
	return authHeader, nil
}

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789#$%&!?"
	letterMax   = len(letterBytes)
)

func randCryptoString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	buffer := make([]byte, n)
	_, err := rand.Read(buffer)
	if err != nil {
		panic(err)
	}
	for i := 0; i < n; i++ {
		sb.WriteByte(letterBytes[buffer[i]%byte(letterMax)])
	}
	return sb.String()
}
