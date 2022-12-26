package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"go.uber.org/zap"
)

type Server struct {
	Issuer  string
	Subject string
	key     jwk.Set
}

func NewServer(ctx context.Context, fileName string) (*Server, error) {
	key, err := jwk.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read in jwk: %w", err)
	}
	return &Server{
		key: key,
	}, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.key)
}

func (s *Server) VerifySessionJWT(token string) (*models.Session, error) {
	return verifySessionJWT(token, s.key)
}

func (s *Server) CreateSessionJWT(sess models.Session) (string, error) {
	// Get top signing key
	key, found := s.key.Key(0)
	if !found {
		return "", fmt.Errorf("failed to get signing key")
	}
	zap.L().Info("Found key 0", zap.String("kid", key.KeyID()))
	// Build a JWT!
	token, err := jwt.NewBuilder().
		Issuer(s.Issuer).
		NotBefore(sess.CreatedAt).
		IssuedAt(time.Now()).
		Expiration(sess.ExpiredAt).
		Subject(s.Subject).
		Claim("sessionID", sess.ID).
		Claim("userID", sess.UserID).
		Build()
	if err != nil {
		return "", fmt.Errorf("failed to build token: %w", err)
	}
	// Sign a JWT!
	signedToken, err := jwt.Sign(token, jwt.WithKey(jwa.EdDSA, key))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return string(signedToken), nil
}
