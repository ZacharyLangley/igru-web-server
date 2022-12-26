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
)

var (
	deletePlantID string
)

func init() {
	deleteCmd.Flags().StringVar(&deletePlantID, "id", "", "ID of an existing plant")
	deleteCmd.MarkFlagRequired("id")
	deleteCmd.MarkFlagRequired("config")
	RootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use: "delete",
	Aliases: []string{
		"rm",
		"del",
	},
	Short:   "Delete an existing plant",
	PreRunE: config.SetupCobraLogger,
	RunE:    deleteGarden,
}

func deleteGarden(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	userClient := gardensv1connect.NewPlantsServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardensv1.DeletePlantRequest{
		Id: deletePlantID,
	})
	_, err := userClient.DeletePlant(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to delete plant: %w", err)
	}
	return nil
}
