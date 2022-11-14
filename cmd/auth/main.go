package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/connect"
	internalcontext "github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/authentication"
	"go.uber.org/zap"
)

type Config struct {
	Database config.Database `mapstructure:"database"`
	GRPC     config.GRPC     `mapstructure:"grpc"`
	Logger   config.Logger   `mapstructure:"logger"`
	GCPeriod time.Duration   `mapstructure:"gcPeriod"`
}

func main() {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		log.Fatalln("failed to parse config", err)
	}
	if err := cfg.Logger.Setup(); err != nil {
		log.Fatalln("failed to setup logger", err)
	}
	baseCtx, done := context.WithCancel(context.Background())
	defer done()
	ctx := internalcontext.New(baseCtx).Named("auth")
	conn, err := connectToDB(ctx, cfg.Database)
	if err != nil {
		ctx.L().Fatal("failed to connect to database", zap.Error(err))
	}
	mux := connect.CreateMux(authentication.New(conn))
	srv := &http.Server{
		Addr:    cfg.GRPC.Address,
		Handler: mux,
	}

	// Start cleanup process
	go func(ctx internalcontext.Context) {
		t := time.NewTicker(cfg.GCPeriod)
		ctx.L().Info("starting gc")
		for range t.C {
			queries := models.New(conn)
			sessions, err := queries.GetExpiredSessions(context.Background())
			if err != nil {
				ctx.L().Error("failed to get expired sessions", zap.Error(err))
				continue
			}
			ctx.L().Debug("deleting old sessions", zap.Int("expiredSessions", len(sessions)))
			for _, session := range sessions {
				if session.ExpiredAt.Before(time.Now()) {
					ctx.L().Debug("deleting session", zap.String("sessionID", session.ID.String()))
					if err := queries.DeleteSession(context.Background(), session.ID); err != nil {
						ctx.L().Error("failed to delete expired session", zap.String("sessionID", session.ID.String()), zap.Error(err))
					}
				}
			}
		}
	}(ctx)
	lc := net.ListenConfig{}
	listener, err := lc.Listen(ctx, "tcp", cfg.GRPC.Address)
	if err != nil {
		ctx.L().Fatal("failed to open listener", zap.String("address", cfg.GRPC.Address))
	}
	ctx.L().Info("listening", zap.String("address", listener.Addr().String()))
	if err := srv.Serve(listener); err != nil {
		ctx.L().Fatal("failed to serve", zap.Error(err))
	}
}

func connectToDB(ctx internalcontext.Context, cfg config.Database) (*sql.DB, error) {
	dsn, err := cfg.DSN()
	if err != nil {
		return nil, err
	}
	backoff := time.Second
	var lastErr error
	var connection *sql.DB
	logger := ctx.L()
	logger.Info("connecting to db")
	for i := 0; i < cfg.MaxRetries; i++ {
		connection, lastErr = database.Open(ctx, dsn, cfg.DisableMigration)
		if lastErr == nil {
			logger.Info("connected to db")
			return connection, nil
		}
		logger.Warn("retrying", zap.Duration("backoff", backoff), zap.Int("attempt", i), zap.Error(lastErr))
		time.Sleep(backoff)
		backoff *= 2
	}
	return nil, fmt.Errorf("failed to connect in time: %w", lastErr)
}
