package ingress

import (
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/garden/v1/gardenv1connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/proxy"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

//go:embed public
var content embed.FS
var webContent fs.FS

type Config struct {
	Metrics config.Metrics `mapstructure:"metrics"`
	Clients struct {
		Authentication config.GRPC `mapstructure:"authentication"`
		Garden         config.GRPC `mapstructure:"garden"`
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

func runServer(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	ctx := context.New(cmd.Context())
	// Setup tracing
	cfg.Metrics.Setup("ingress")
	// Create and populate Mux
	zap.L().Info("Setting up router")
	r := mux.NewRouter()
	r.MethodNotAllowedHandler = errorHandler("method not allowed")
	r.NotFoundHandler = errorHandler("not found")
	r.Use(loggingMiddleware)
	// Attach services
	if err := proxy.RegisterProxy(r, cfg.Clients.Authentication, authenticationv1connect.UserServiceName); err != nil {
		return fmt.Errorf("failed to register authentication proxy")
	}
	if err := proxy.RegisterProxy(r, cfg.Clients.Authentication, authenticationv1connect.GroupServiceName); err != nil {
		return fmt.Errorf("failed to register authentication proxy")
	}
	if err := proxy.RegisterProxy(r, cfg.Clients.Authentication, authenticationv1connect.SessionServiceName); err != nil {
		return fmt.Errorf("failed to register authentication proxy")
	}
	if err := proxy.RegisterProxy(r, cfg.Clients.Garden, gardenv1connect.GardenServiceName); err != nil {
		return fmt.Errorf("failed to register gardens proxy")
	}
	if err := proxy.RegisterProxy(r, cfg.Clients.Garden, gardenv1connect.PlantServiceName); err != nil {
		return fmt.Errorf("failed to register gardens proxy")
	}
	if err := proxy.RegisterProxy(r, cfg.Clients.Garden, gardenv1connect.StrainServiceName); err != nil {
		return fmt.Errorf("failed to register gardens proxy")
	}
	if err := proxy.RegisterProxy(r, cfg.Clients.Garden, gardenv1connect.RecipeServiceName); err != nil {
		return fmt.Errorf("failed to register gardens proxy")
	}

	// Attach embedded frontend
	webContent, err := fs.Sub(content, "public")
	if err != nil {
		return fmt.Errorf("failed to find web content")
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
		r.Methods("GET").Handler(proxy.HTTP{URL: proxyURL})
	}
	// Create TCP listener service
	lc := net.ListenConfig{}
	listener, err := lc.Listen(ctx, "tcp", cfg.GRPC.Address)
	if err != nil {
		return fmt.Errorf("failed to open listener to %q: %w", cfg.GRPC.Address, err)
	}
	ctx.L().Info("listening", zap.String("address", listener.Addr().String()))
	// Start serving HTTP services
	if err := http.Serve(listener, r); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}
	return nil
}

func errorHandler(message string) http.Handler {
	logger := zap.L()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		logger.Debug(message,
			zap.String("requestURI", r.RequestURI))
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	logger := zap.L()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		logger.Debug("processing request",
			zap.String("requestURI", r.RequestURI))
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
