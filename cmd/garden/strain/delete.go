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
)

var (
	deleteStrainID string
)

func init() {
	deleteCmd.Flags().StringVar(&deleteStrainID, "id", "", "ID of an existing strain")
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
	Short:   "Delete an existing strain",
	PreRunE: config.SetupCobraLogger,
	RunE:    deleteStrain,
}

func deleteStrain(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	userClient := gardensv1connect.NewStrainsServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardensv1.DeleteStrainRequest{
		Id: deleteStrainID,
	})
	_, err := userClient.DeleteStrain(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to delete strain: %w", err)
	}
	return nil
}
