package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/authentication"
)

type Config struct {
	Database config.Database `mapstructure:"database"`
	GRPC     config.GRPC     `mapstructure:"grpc"`
	GCPeriod time.Duration   `mapstructure:"gcPeriod"`
}

func main() {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		log.Fatalln(err)
	}
	log.Println("Starting authentication service")

	conn, err := connectToDB(context.Background(), cfg.Database)
	if err != nil {
		log.Fatalln("Can't connect to Postgres!", err)
	}

	mux := connect.CreateMux(authentication.New(conn))
	srv := &http.Server{
		Addr:    cfg.GRPC.Address,
		Handler: mux,
	}

	// Start cleanup process
	go func(ctx context.Context) {
		t := time.NewTicker(cfg.GCPeriod)
		for range t.C {
			queries := models.New(conn)
			sessions, err := queries.GetExpiredSessions(context.Background())
			if err != nil {
				log.Println("Failed to get expired sessions", err)
			}
			log.Println("Clearing old sessions", len(sessions))
			for _, session := range sessions {
				if session.ExpiredAt.Before(time.Now()) {
					log.Println("Deleting", session.ID)
					if err := queries.DeleteSession(context.Background(), session.ID); err != nil {
						log.Println("Failed to delete expired sessions")
					}
				}
			}
		}
	}(context.Background())
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func connectToDB(ctx context.Context, cfg config.Database) (*sql.DB, error) {
	dsn, err := cfg.DSN()
	if err != nil {
		return nil, err
	}
	backoff := time.Second
	var lastErr error
	var connection *sql.DB
	for i := 0; i < cfg.MaxRetries; i++ {
		connection, lastErr = database.Open(ctx, dsn)
		if err == nil {
			return connection, nil
		}
		time.Sleep(backoff)
		backoff *= 2
	}
	return nil, fmt.Errorf("failed to connect in time: %w", lastErr)
}
