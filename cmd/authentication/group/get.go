package group

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
	config.Must(getCmd.MarkFlagRequired("config"))
	RootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use: "get",
	Aliases: []string{
		"list",
	},
	Short:   "Get all existing user group",
	PreRunE: config.SetupCobraLogger,
	RunE:    getGroup,
}

func getGroup(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	userClient := authenticationv1connect.NewGroupServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&authenticationv1.GetGroupsRequest{})
	resp, err := userClient.GetGroups(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get groups: %w", err)
	}
	for _, group := range resp.Msg.Groups {
		zap.L().Info("Found group", zap.Any("group", group))
	}
	return nil
}
