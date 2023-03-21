package garden

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
	deleteGardenID string
)

func init() {
	deleteCmd.Flags().StringVar(&deleteGardenID, "id", "", "ID of an existing garden")
	config.Must(deleteCmd.MarkFlagRequired("id"))
	RootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use: "delete",
	Aliases: []string{
		"rm",
		"del",
	},
	Short:   "Delete an existing garden",
	PreRunE: config.SetupCobraLogger,
	RunE:    deleteGarden,
}

func deleteGarden(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	gardenClient := gardenv1connect.NewGardenServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardenv1.DeleteGardenRequest{
		Id: deleteGardenID,
	})
	_, err := gardenClient.DeleteGarden(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to delete garden: %w", err)
	}
	return nil
}
