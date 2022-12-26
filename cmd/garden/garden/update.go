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
)

var (
	updateID            string
	updateName          string
	updateComment       string
	updateLocation      string
	updateGrowType      string
	updateGrowSize      string
	updateContainerSize string
	updateTags          string
)

func init() {
	updateCmd.Flags().StringVar(&updateID, "id", "", "ID of an existing garden")
	updateCmd.Flags().StringVar(&updateName, "name", "", "Name of an existing garden")
	updateCmd.Flags().StringVar(&updateComment, "comment", "", "Comment of an existing garden")
	updateCmd.Flags().StringVar(&updateLocation, "location", "", "Location of an existing garden")
	updateCmd.Flags().StringVar(&updateGrowType, "growType", "", "Grow type of an existing garden")
	updateCmd.Flags().StringVar(&updateGrowSize, "growSize", "", "Grow size of an existing garden")
	updateCmd.Flags().StringVar(&updateContainerSize, "containerSize", "", "ContainerSize of an existing garden")
	updateCmd.Flags().StringVar(&updateTags, "tags", "", "Tags of an existing garden")
	updateCmd.MarkFlagRequired("name")
	updateCmd.MarkFlagRequired("config")
	RootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update an existing garden",
	PreRunE: config.SetupCobraLogger,
	RunE:    updateGarden,
}

func updateGarden(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	userClient := gardensv1connect.NewGardensServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardensv1.UpdateGardenRequest{
		Id:            updateID,
		Name:          updateName,
		Comment:       updateComment,
		Location:      updateLocation,
		GrowType:      updateGrowType,
		GrowSize:      updateGrowSize,
		ContainerSize: updateContainerSize,
		Tags:          updateTags,
	})
	_, err := userClient.UpdateGarden(ctx, req)
	if err != nil {
		return fmt.Errorf("Failed to update garden: %w", err)
	}
	return nil
}
