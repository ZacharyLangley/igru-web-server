package strain

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
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use: "list",
	Aliases: []string{
		"l",
	},
	Short:   "Get all existing strains",
	PreRunE: config.SetupCobraLogger,
	RunE:    getStrains,
}

func getStrains(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	recipeClient := gardensv1connect.NewStrainsServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardensv1.GetStrainsRequest{})
	resp, err := recipeClient.GetStrains(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get strains: %w", err)
	}
	if resp.Msg.Strains != nil {
		zap.L().Info("Found strains", zap.Any("strains", resp.Msg.Strains))
	}
	return nil
}
