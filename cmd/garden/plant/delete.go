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
)

var (
	deletePlantID string
)

func init() {
	deleteCmd.Flags().StringVar(&deletePlantID, "id", "", "ID of an existing plant")
	config.Must(deleteCmd.MarkFlagRequired("id"))
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
	RunE:    deletePlant,
}

func deletePlant(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	plantClient := gardenv1connect.NewPlantServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardenv1.DeletePlantRequest{
		Id: deletePlantID,
	})
	_, err := plantClient.DeletePlant(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to delete plant: %w", err)
	}
	return nil
}
