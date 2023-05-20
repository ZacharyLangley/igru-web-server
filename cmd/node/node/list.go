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

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use: "list",
	Aliases: []string{
		"l",
	},
	Short:   "Get all existing nodes",
	PreRunE: config.SetupCobraLogger,
	RunE:    getNodes,
}

func getNodes(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	nodeClient := nodev1connect.NewNodeServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&nodev1.GetNodesRequest{})
	resp, err := nodeClient.GetNodes(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get nodes: %w", err)
	}
	if resp.Msg.Nodes != nil {
		zap.L().Info("Found nodes", zap.Any("nodes", resp.Msg.Nodes))
	}
	return nil
}
