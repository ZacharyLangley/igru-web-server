package ingress

import (
	"fmt"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/broker"
	"github.com/spf13/cobra"
)

type Config struct {
	GRPC    config.GRPC    `mapstructure:"grpc"`
	Metrics config.Metrics `mapstructure:"metrics"`
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
	// Start serving
	return connect.ServeMux(ctx, cfg.GRPC, broker.New())
}
