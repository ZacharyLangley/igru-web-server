package node

import (
	"fmt"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	nodev1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/node/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/node/v1/nodev1connect"
	"github.com/bufbuild/connect-go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	getNodeMacAddress string
)

func init() {
	getCmd.Flags().StringVar(&getNodeMacAddress, "mac_address", "", "MAC address of an existing node")
	config.Must(getCmd.MarkFlagRequired("mac_address"))
	RootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use: "get",
	Aliases: []string{
		"g",
	},
	Short:   "Get an existing node",
	PreRunE: config.SetupCobraLogger,
	RunE:    getNode,
}

func getNode(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	nodeClient := nodev1connect.NewNodeServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&nodev1.GetNodeRequest{
		MacAddress: getNodeMacAddress,
	})
	resp, err := nodeClient.GetNode(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get node: %w", err)
	}
	if resp.Msg.Node != nil {
		zap.L().Info("Found node", zap.Any("node", resp.Msg.Node))
	}
	return nil
}
