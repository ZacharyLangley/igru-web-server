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
)

var (
	addGroupID string
	addUserID  string
)

func init() {
	addCmd.Flags().StringVar(&addGroupID, "groupID", "", "ID of an existing user group")
	addCmd.MarkFlagRequired("groupID")
	addCmd.Flags().StringVar(&addUserID, "userID", "", "ID of an existing user")
	addCmd.MarkFlagRequired("userID")
	addCmd.MarkFlagRequired("config")
	RootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use: "add",
	Aliases: []string{
		"new",
		"create",
	},
	Short:   "Add a new member to a user group",
	PreRunE: config.SetupCobraLogger,
	RunE:    createMember,
}

func createMember(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	userClient := authenticationv1connect.NewGroupServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&authenticationv1.AddGroupMemberRequest{
		GroupId: addGroupID,
		UserId:  addUserID,
	})
	_, err := userClient.AddGroupMember(ctx, req)
	if err != nil {
		return fmt.Errorf("Failed to add group member: %w", err)
	}
	return nil
}
