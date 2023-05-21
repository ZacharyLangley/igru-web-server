package ingress

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/ZacharyLangley/igru-web-server/pkg/middleware"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/garden/v1/gardenv1connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/node/v1/nodev1connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/proxy"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.uber.org/zap"
)

//go:embed public
var content embed.FS

type Config struct {
	Metrics config.Metrics `mapstructure:"metrics"`
	Clients struct {
		Authentication config.GRPC `mapstructure:"authentication"`
		Garden         config.GRPC `mapstructure:"garden"`
		Node           config.GRPC `mapstructure:"node"`
	} `mapstructure:"clients"`
	WebProxyAddress string      `mapstructure:"webProxyAddress"`
	GRPC            config.GRPC `mapstructure:"grpc"`
}

var serveCmd = &cobra.Command{
	Use:     "serve",
	Short:   "Start the server",
	PreRunE: config.SetupCobraLogger,
	RunE:    runServer,
}

var errWebContentNotFound = errors.New("web content not found")

func runServer(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	ctx := context.New(cmd.Context())
	// Setup tracing
	if err := cfg.Metrics.Setup("ingress"); err != nil {
		return fmt.Errorf("failed to setup metrics: %w", err)
	}
	// Create and populate Mux
	zap.L().Info("Setting up router")
	r := mux.NewRouter()
	r.Use(otelmux.Middleware("ingress"))
	r.Use(middleware.HTTPLoggingMiddleware)
	// Attach services
	if err := proxy.RegisterProxy(r, cfg.Clients.Authentication, authenticationv1connect.UserServiceName); err != nil {
		return fmt.Errorf("failed to register authentication proxy: %w", err)
	}
	if err := proxy.RegisterProxy(r, cfg.Clients.Authentication, authenticationv1connect.GroupServiceName); err != nil {
		return fmt.Errorf("failed to register authentication proxy: %w", err)
	}
	if err := proxy.RegisterProxy(r, cfg.Clients.Authentication, authenticationv1connect.SessionServiceName); err != nil {
		return fmt.Errorf("failed to register authentication proxy: %w", err)
	}
	if err := proxy.RegisterProxy(r, cfg.Clients.Garden, gardenv1connect.GardenServiceName); err != nil {
		return fmt.Errorf("failed to register gardens proxy: %w", err)
	}
	if err := proxy.RegisterProxy(r, cfg.Clients.Garden, gardenv1connect.PlantServiceName); err != nil {
		return fmt.Errorf("failed to register gardens proxy: %w", err)
	}
	if err := proxy.RegisterProxy(r, cfg.Clients.Garden, gardenv1connect.StrainServiceName); err != nil {
		return fmt.Errorf("failed to register gardens proxy: %w", err)
	}
	if err := proxy.RegisterProxy(r, cfg.Clients.Garden, gardenv1connect.RecipeServiceName); err != nil {
		return fmt.Errorf("failed to register gardens proxy: %w", err)
	}
	if err := proxy.RegisterProxy(r, cfg.Clients.Node, nodev1connect.NodeServiceName); err != nil {
		return fmt.Errorf("failed to register node proxy: %w", err)
	}

	// Attach embedded frontend
	webContent, err := fs.Sub(content, "public")
	if err != nil {
		return errWebContentNotFound
	}
	if cfg.WebProxyAddress == "" {
		zap.L().Info("Using FS server")
		r.Methods("GET").Handler(http.FileServer(http.FS(webContent)))
	} else {
		// Proxy frontend if required
		zap.L().Info("Using HTTP proxy")
		proxyURL, err := url.Parse("http://" + cfg.WebProxyAddress)
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed parse proxy url", err)
			os.Exit(-1)
		}
		r.Methods("GET").Handler(proxy.NewTunnel("http", proxyURL))
	}
	// Create TCP listener service
	lc := net.ListenConfig{}
	listener, err := lc.Listen(ctx, "tcp", cfg.GRPC.Address)
	if err != nil {
		return fmt.Errorf("failed to open listener to %q: %w", cfg.GRPC.Address, err)
	}
	ctx.L().Info("listening", zap.String("address", listener.Addr().String()))
	// Start serving HTTP services
	httpServer := http.Server{
		Handler:      r,
		ReadTimeout:  requestTimeout,
		WriteTimeout: requestTimeout,
		IdleTimeout:  requestTimeout,
	}
	if err := httpServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}
	return nil
}

var requestTimeout = time.Second * 30
