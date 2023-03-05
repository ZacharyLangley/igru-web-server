package garden

import (
	"fmt"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/auth"
	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/garden"
	"github.com/spf13/cobra"
)

type Config struct {
	Database       config.Database `mapstructure:"database"`
	GRPC           config.GRPC     `mapstructure:"grpc"`
	Authentication config.GRPC     `mapstructure:"authentication"`
	Logger         config.Logger   `mapstructure:"logger"`
	Metrics        config.Metrics  `mapstructure:"metrics"`
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
	if err := cfg.Metrics.Setup("garden"); err != nil {
		return fmt.Errorf("failed to setup metrics: %w", err)
	}
	// Connect to DB
	conn, err := database.Open(ctx, cfg.Database)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	// Create auth client
	checker := &auth.Checker{
		SessionServiceClient: authenticationv1connect.NewSessionServiceClient(
			http.DefaultClient,
			cfg.Authentication.Address,
		),
	}
	// Start serving
	return connect.ServeMux(ctx, cfg.GRPC, garden.New(conn, checker))
}
