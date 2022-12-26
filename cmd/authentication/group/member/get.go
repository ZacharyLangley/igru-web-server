package member

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

var getGroupID string

func init() {
	getCmd.Flags().StringVar(&getGroupID, "groupID", "", "ID of an existing user group")
	getCmd.MarkFlagRequired("groupID")
	getCmd.MarkFlagRequired("config")
	RootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use: "get",
	Aliases: []string{
		"list",
	},
	Short:   "Get all user group members",
	PreRunE: config.SetupCobraLogger,
	RunE:    getMembers,
}

func getMembers(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	userClient := authenticationv1connect.NewGroupServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&authenticationv1.GetGroupMembersRequest{
		GroupId: getGroupID,
	})
	resp, err := userClient.GetGroupMembers(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get groups: %w", err)
	}
	for _, member := range resp.Msg.Members {
		zap.L().Info("Found member", zap.Any("member", member))
	}
	return nil
}
