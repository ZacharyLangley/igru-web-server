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
	removeGroupID string
	removeUserID  string
)

func init() {
	removeCmd.Flags().StringVar(&removeGroupID, "groupID", "", "ID of an existing user group")
	config.Must(removeCmd.MarkFlagRequired("groupID"))
	removeCmd.Flags().StringVar(&removeUserID, "userID", "", "ID of an existing user")
	config.Must(removeCmd.MarkFlagRequired("userID"))
	RootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use: "remove",
	Aliases: []string{
		"rm",
		"del",
		"delete",
	},
	Short:   "Remove a member from a user group",
	PreRunE: config.SetupCobraLogger,
	RunE:    removeMember,
}

func removeMember(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	userClient := authenticationv1connect.NewGroupServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&authenticationv1.RemoveGroupMemberRequest{
		UserId:  removeUserID,
		GroupId: removeGroupID,
	})
	_, err := userClient.RemoveGroupMember(ctx, req)
	if err != nil {
		return fmt.Errorf("Failed to remove group member: %w", err)
	}
	return nil
}
