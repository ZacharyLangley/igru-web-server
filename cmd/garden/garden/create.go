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
	"go.uber.org/zap"
)

var (
	createName          string
	createComment       string
	createLocation      string
	createGrowType      string
	createGrowStyle     string
	createGrowSize      string
	createContainerSize string
	createTags          string
)

func init() {
	createCmd.Flags().StringVar(&createName, "name", "", "Name of a new garden")
	createCmd.Flags().StringVar(&createComment, "comment", "", "Comment of a new garden")
	createCmd.Flags().StringVar(&createLocation, "location", "", "Location of a new garden")
	createCmd.Flags().StringVar(&createGrowType, "growType", "", "Grow type of a new garden")
	createCmd.Flags().StringVar(&createGrowStyle, "growStyle", "", "Grow style of a new garden")
	createCmd.Flags().StringVar(&createGrowSize, "growSize", "", "Grow size of a new garden")
	createCmd.Flags().StringVar(&createContainerSize, "containerSize", "", "ContainerSize of a new garden")
	createCmd.Flags().StringVar(&createTags, "tags", "", "Tags of a new garden")
	createCmd.MarkFlagRequired("name")
	RootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use: "create",
	Aliases: []string{
		"new",
	},
	Short:   "Create a new garden",
	PreRunE: config.SetupCobraLogger,
	RunE:    createGarden,
}

func createGarden(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	gardenClient := gardensv1connect.NewGardensServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardensv1.CreateGardenRequest{
		Name:          createName,
		Comment:       createComment,
		Location:      createLocation,
		GrowType:      createGrowType,
		GrowStyle:     createGrowStyle,
		GrowSize:      createGrowSize,
		ContainerSize: createContainerSize,
		Tags:          createTags,
	})
	resp, err := gardenClient.CreateGarden(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to add garden: %w", err)
	}
	zap.L().Info("Created garden", zap.Any("garden", resp.Msg.Garden))
	return nil
}
