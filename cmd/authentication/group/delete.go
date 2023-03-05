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
)

var (
	deleteGroupID string
)

func init() {
	deleteCmd.Flags().StringVar(&deleteGroupID, "id", "", "ID of an existing group")
	config.Must(deleteCmd.MarkFlagRequired("id"))
	config.Must(deleteCmd.MarkFlagRequired("config"))
	RootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use: "delete",
	Aliases: []string{
		"rm",
		"del",
	},
	Short:   "Delete an existing user group",
	PreRunE: config.SetupCobraLogger,
	RunE:    deleteGroup,
}

func deleteGroup(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	userClient := authenticationv1connect.NewGroupServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&authenticationv1.DeleteGroupRequest{
		Id: deleteGroupID,
	})
	_, err := userClient.DeleteGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to delete group: %w", err)
	}
	return nil
}
