package user

import (
	"fmt"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	authenticationv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/bufbuild/connect-go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	getCmd.MarkFlagRequired("config")
	RootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use: "get",
	Aliases: []string{
		"list",
	},
	Short:   "Get existing users",
	PreRunE: config.SetupCobraLogger,
	RunE:    getUser,
}

func getUser(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	userClient := authenticationv1connect.NewUserServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&authenticationv1.GetUsersRequest{})
	resp, err := userClient.GetUsers(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}
	for _, user := range resp.Msg.Users {
		zap.L().Info("Found user", zap.Any("user", user))
	}
	return nil
}
