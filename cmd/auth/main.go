package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/authentication"
	"github.com/go-redis/redis/v8"
)

type Config struct {
	Database config.SQL    `mapstructure:"database"`
	GRPC     config.GRPC   `mapstructure:"grpc"`
	GCPeriod time.Duration `mapstructure:"gcPeriod"`
	Cache    config.Cache  `mapstructure:"cache"`
}

func main() {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		log.Fatalln(err)
	}
	log.Println("Starting authentication service")
	ctx, done := context.WithCancel(context.Background())
	defer done()

	cache, err := connectToCache(ctx, cfg.Cache)
	if err != nil {
		log.Fatalln("Can't connect to Redis!", err)
	}

	conn, err := connectToDB(ctx, cfg.Database)
	if err != nil {
		log.Fatalln("Can't connect to Postgres!", err)
	}

	mux := connect.CreateMux(authentication.New(conn, cache))
	srv := &http.Server{
		Addr:    cfg.GRPC.Address,
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func connectToDB(ctx context.Context, cfg config.SQL) (*sql.DB, error) {
	dsn, err := cfg.DSN()
	if err != nil {
		return nil, err
	}
	backoff := time.Second
	var lastErr error
	var connection *sql.DB
	for i := 0; i < cfg.MaxRetries; i++ {
		connection, lastErr = database.Open(ctx, dsn, cfg.DisableMigration)
		if err == nil && connection != nil {
			return connection, nil
		}
		time.Sleep(backoff)
		backoff *= 2
	}
	return nil, fmt.Errorf("failed to connect in time: %w", lastErr)
}

func connectToCache(ctx context.Context, cfg config.Cache) (*redis.Client, error) {
	cli := redis.NewClient(cfg.Options())
	if cli == nil {
		return nil, errors.New("failed to connect to redis")
	}
	return cli, nil
}
