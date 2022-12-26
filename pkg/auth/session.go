package auth

import (
	"context"
	"fmt"
	"time"

	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type Service struct {
	issuer  string
	subject string
	url     string
	cache   *jwk.Cache
}

func NewService(
	rootCtx context.Context,
	refreshPeriod time.Duration,
	url string,
	issuer string,
	subject string) (*Service, error) {
	ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()
	var s Service
	s.cache = jwk.NewCache(ctx)
	// Register and refresh URL
	if err := s.cache.Register(url, jwk.WithMinRefreshInterval(refreshPeriod)); err != nil {
		return nil, fmt.Errorf("failed to register %q: %w", url, err)
	}
	_, err := s.cache.Refresh(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh %q: %w", url, err)
	}
	s.url = url
	s.issuer = issuer
	s.subject = subject
	return &s, nil
}

func verifySessionJWT(token string, keySet jwk.Set) (*models.Session, error) {
	parsedToken, err := jwt.ParseString(token, jwt.WithKeySet(keySet))
	if err != nil {
		return nil, err
	}
	var ok bool
	var sess models.Session
	rawSessionID, ok := parsedToken.Get("sessionID")
	if err != nil {
		return nil, fmt.Errorf("failed to find session id")
	}
	sess.ID, ok = rawSessionID.(uuid.UUID)
	if !ok {
		return nil, fmt.Errorf("failed to parse sessions id")
	}
	rawUserID, ok := parsedToken.Get("userID")
	if err != nil {
		return nil, fmt.Errorf("failed to find user id")
	}
	sess.UserID, ok = rawUserID.(uuid.UUID)
	if !ok {
		return nil, fmt.Errorf("failed to parse sesusersions id")
	}
	sess.ExpiredAt = parsedToken.Expiration()
	sess.CreatedAt = parsedToken.NotBefore()
	return &sess, nil
}

func (s *Service) VerifySessionJWT(ctx context.Context, token string) (*models.Session, error) {
	keySet, err := s.cache.Get(ctx, s.url)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve key set: %w", err)
	}
	return verifySessionJWT(token, keySet)
}
