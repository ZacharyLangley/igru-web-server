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
	"go.uber.org/zap"
)

var (
	updateID            string
	updateName          string
	updateComment       string
	updateLocation      string
	updateGrowType      string
	updateGrowStyle     string
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
	updateCmd.Flags().StringVar(&updateGrowStyle, "growStyle", "", "Grow style of an existing garden")
	updateCmd.Flags().StringVar(&updateGrowSize, "growSize", "", "Grow size of an existing garden")
	updateCmd.Flags().StringVar(&updateContainerSize, "containerSize", "", "ContainerSize of an existing garden")
	updateCmd.Flags().StringVar(&updateTags, "tags", "", "Tags of an existing garden")
	config.Must(updateCmd.MarkFlagRequired("name"))
	config.Must(updateCmd.MarkFlagRequired("config"))
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
	gardenClient := gardenv1connect.NewGardenServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardenv1.UpdateGardenRequest{
		Id:            updateID,
		Name:          updateName,
		Comment:       updateComment,
		Location:      updateLocation,
		GrowType:      updateGrowType,
		GrowStyle:     updateGrowStyle,
		GrowSize:      updateGrowSize,
		ContainerSize: updateContainerSize,
		Tags:          updateTags,
	})
	resp, err := gardenClient.UpdateGarden(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to update garden: %w", err)
	}
	if resp.Msg.Garden != nil {
		zap.L().Info("updated garden", zap.Any("garden", resp.Msg.Garden))
	}
	return nil
}
