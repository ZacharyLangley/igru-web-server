package garden

import (
	"fmt"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/gardens"
	"github.com/spf13/cobra"
)

type Config struct {
	Database config.Database `mapstructure:"database"`
	GRPC     config.GRPC     `mapstructure:"grpc"`
	Logger   config.Logger   `mapstructure:"logger"`
	Metrics  config.Metrics  `mapstructure:"metrics"`
	GCPeriod time.Duration   `mapstructure:"gcPeriod"`
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
	cfg.Metrics.Setup()
	// Connect to DB
	conn, err := database.Open(ctx, cfg.Database)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	// Start serving
	return connect.ServeMux(ctx, cfg.GRPC, gardens.New(conn))
}
