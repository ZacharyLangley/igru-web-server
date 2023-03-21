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

var (
	createName string
)

func init() {
	createCmd.Flags().StringVar(&createName, "name", "", "Name of a new group")
	config.Must(createCmd.MarkFlagRequired("name"))
	RootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use: "create",
	Aliases: []string{
		"new",
	},
	Short:   "Create a new user group",
	PreRunE: config.SetupCobraLogger,
	RunE:    createGroup,
}

func createGroup(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	userClient := authenticationv1connect.NewGroupServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&authenticationv1.CreateGroupRequest{
		Name: createName,
	})
	resp, err := userClient.CreateGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("Failed to add group: %w", err)
	}
	zap.L().Info("Created group", zap.Any("group", resp.Msg.Group))
	return nil
}
