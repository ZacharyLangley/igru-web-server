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
	getGardenID string
)

func init() {
	config.Must(getCmd.MarkFlagRequired("config"))
	getCmd.Flags().StringVar(&getGardenID, "id", "", "id of an existing garden")
	config.Must(getCmd.MarkFlagRequired("id"))
	RootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use: "get",
	Aliases: []string{
		"g",
	},
	Short:   "Get an existing garden",
	PreRunE: config.SetupCobraLogger,
	RunE:    getGroup,
}

func getGroup(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	gardenClient := gardenv1connect.NewGardenServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardenv1.GetGardenRequest{
		Id: getGardenID,
	})
	resp, err := gardenClient.GetGarden(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get garden: %w", err)
	}
	if resp.Msg.Garden != nil {
		zap.L().Info("Found garden", zap.Any("garden", resp.Msg.Garden))
	}
	return nil
}
