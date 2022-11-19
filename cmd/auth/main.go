package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/connect"
	internalcontext "github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/authentication"
	"go.uber.org/zap"
)

type Config struct {
	Database config.Database `mapstructure:"database"`
	GRPC     config.GRPC     `mapstructure:"grpc"`
	Logger   config.Logger   `mapstructure:"logger"`
	Metrics  config.Metrics  `mapstructure:"metrics"`
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
	// Setup tracing
	cfg.Metrics.Setup()
	// Connect to DB
	conn, err := database.Open(ctx, cfg.Database)
	if err != nil {
		ctx.L().Fatal("failed to connect to database", zap.Error(err))
	}
	mux := connect.CreateMux(authentication.New(conn))
	srv := &http.Server{
		Addr:    cfg.GRPC.Address,
		Handler: mux,
	}

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
