package strain

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
	deleteStrainID string
)

func init() {
	deleteCmd.Flags().StringVar(&deleteStrainID, "id", "", "ID of an existing strain")
	config.Must(deleteCmd.MarkFlagRequired("id"))
	RootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use: "delete",
	Aliases: []string{
		"rm",
		"del",
	},
	Short:   "Delete an existing strain",
	PreRunE: config.SetupCobraLogger,
	RunE:    deleteStrain,
}

func deleteStrain(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	recipeClient := gardenv1connect.NewStrainServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardenv1.DeleteStrainRequest{
		Id: deleteStrainID,
	})
	_, err := recipeClient.DeleteStrain(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to delete strain: %w", err)
	}
	return nil
}
