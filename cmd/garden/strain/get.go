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

var (
	getStrainID string
)

func init() {
	getCmd.MarkFlagRequired("config")
	getCmd.Flags().StringVar(&getStrainID, "id", "", "id of an existing strain")
	getCmd.MarkFlagRequired("id")
	RootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use: "get",
	Aliases: []string{
		"g",
	},
	Short:   "Get an existing strain",
	PreRunE: config.SetupCobraLogger,
	RunE:    getStrain,
}

func getStrain(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	recipeClient := gardensv1connect.NewStrainsServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardensv1.GetStrainRequest{
		Id: getStrainID,
	})
	resp, err := recipeClient.GetStrain(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get strains: %w", err)
	}
	if resp.Msg.Strain != nil {
		zap.L().Info("Found strain", zap.Any("strain", resp.Msg.Strain))
	}
	return nil
}
