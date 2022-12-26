package garden

import (
	"fmt"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	gardensv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/gardens/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/gardens/v1/gardensv1connect"
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
	Short:   "Get all existing user group",
	PreRunE: config.SetupCobraLogger,
	RunE:    getGroup,
}

func getGroup(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	userClient := gardensv1connect.NewGardensServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardensv1.GetGardensRequest{})
	resp, err := userClient.GetGardens(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get gardens: %w", err)
	}
	for _, group := range resp.Msg.Gardens {
		zap.L().Info("Found garden", zap.Any("garden", group))
	}
	return nil
}
