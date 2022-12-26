package authentication

import (
	"fmt"

	"github.com/ZacharyLangley/igru-web-server/pkg/auth"
	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/authentication"
	"github.com/spf13/cobra"
)

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
	// Open JWK server
	authServer, err := auth.NewServer(ctx, cfg.AuthServer.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to setup JWK service: %w", err)
	}
	// Start serving
	return connect.ServeMux(ctx, cfg.GRPC, authentication.New(conn, authServer))
}
