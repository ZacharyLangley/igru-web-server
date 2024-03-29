package plant

import (
	"fmt"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	gardenv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/garden/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/garden/v1/gardenv1connect"
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
	Short:   "Get all existing plants",
	PreRunE: config.SetupCobraLogger,
	RunE:    getPlants,
}

func getPlants(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	plantClient := gardenv1connect.NewPlantServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardenv1.GetPlantsRequest{})
	resp, err := plantClient.GetPlants(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get plants: %w", err)
	}
	if resp.Msg.Plants != nil {
		zap.L().Info("Found plants", zap.Any("plants", resp.Msg.Plants))
	}
	return nil
}
