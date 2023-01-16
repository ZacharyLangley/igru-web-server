package plant

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

var (
	getPlantID string
)

func init() {
	getCmd.MarkFlagRequired("config")
	getCmd.Flags().StringVar(&getPlantID, "id", "", "id of an existing plant")
	getCmd.MarkFlagRequired("id")
	RootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use: "get",
	Aliases: []string{
		"g",
	},
	Short:   "Get an existing plant",
	PreRunE: config.SetupCobraLogger,
	RunE:    getPlant,
}

func getPlant(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	plantClient := gardensv1connect.NewPlantsServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardensv1.GetPlantRequest{
		Id: getPlantID,
	})
	resp, err := plantClient.GetPlant(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get plants: %w", err)
	}
	if resp.Msg.Plant != nil {
		zap.L().Info("Found plant", zap.Any("plant", resp.Msg.Plant))
	}
	return nil
}
