package node

import (
	"fmt"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/auth"
	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/node"
	connect_go "github.com/bufbuild/connect-go"
	otelconnect "github.com/bufbuild/connect-opentelemetry-go"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel/propagation"
)

func runServer(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	ctx := context.New(cmd.Context())
	// Setup tracing
	if err := cfg.Metrics.Setup("node"); err != nil {
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
			connect_go.WithInterceptors(
				otelconnect.NewInterceptor(
					otelconnect.WithTrustRemote(),
					otelconnect.WithPropagator(propagation.TraceContext{}),
				),
			),
		),
	}
	service := node.New(conn, checker)
	// Start serving
	if err := connect.ServeMux(ctx, cfg.GRPC, service); err != nil {
		return fmt.Errorf("failed to server: %w", err)
	}
	return nil
}
