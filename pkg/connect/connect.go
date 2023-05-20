package connect

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"go.uber.org/zap"
)

var serverTimeout = time.Second * 30

type Service interface {
	Register(*http.ServeMux)
}

func ServeMux(ctx context.Context, grpcConfig config.GRPC, services ...Service) error {
	// Create and populate MUX
	mux := http.NewServeMux()
	for _, svc := range services {
		svc.Register(mux)
	}
	mux.HandleFunc("*", func(w http.ResponseWriter, r *http.Request) {
		zap.L().Info("Got something")
	})
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	// Create HTTP server
	srv := &http.Server{
		Addr:         grpcConfig.Address,
		Handler:      mux,
		ReadTimeout:  serverTimeout,
		WriteTimeout: serverTimeout,
		IdleTimeout:  serverTimeout,
	}
	// Create TCP listener service
	lc := net.ListenConfig{}
	listener, err := lc.Listen(ctx, "tcp", grpcConfig.Address)
	if err != nil {
		return fmt.Errorf("failed to open listener to %q: %w", grpcConfig.Address, err)
	}
	ctx.L().Info("listening", zap.String("address", listener.Addr().String()))
	// Start serving HTTP services
	if err := srv.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}
	return nil
}
