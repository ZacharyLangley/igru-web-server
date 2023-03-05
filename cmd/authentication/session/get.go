package session

import (
	"fmt"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	authenticationv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/authentication"
	"github.com/bufbuild/connect-go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	config.Must(getCmd.MarkFlagRequired("config"))
	RootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use: "get",
	Aliases: []string{
		"list",
	},
	Short:   "Get existing sessions",
	PreRunE: config.SetupCobraLogger,
	RunE:    getSessions,
}

func getSessions(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	clientConfig, err := config.NewClient()
	if err != nil {
		return fmt.Errorf("failed to open new client config: %w", err)
	}
	userClient := authenticationv1connect.NewSessionServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&authenticationv1.GetSessionsRequest{})
	token, err := clientConfig.Get()
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}
	authentication.AddSessionToken(req.Header(), token)
	resp, err := userClient.GetSessions(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get sessions: %w", err)
	}
	for _, session := range resp.Msg.Sessions {
		zap.L().Info("Found session", zap.Any("session", session))
	}
	return nil
}
