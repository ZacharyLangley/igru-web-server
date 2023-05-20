package node

import (
	"github.com/ZacharyLangley/igru-web-server/cmd/node/node"
	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "node",
	Aliases: []string{
		"node",
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
	RootCmd.AddCommand(node.RootCmd)
}

type Config struct {
	Database       config.Database `mapstructure:"database"`
	GRPC           config.GRPC     `mapstructure:"grpc"`
	Authentication config.GRPC     `mapstructure:"authentication"`
	Logger         config.Logger   `mapstructure:"logger"`
	Metrics        config.Metrics  `mapstructure:"metrics"`
}

var serveCmd = &cobra.Command{
	Use:     "serve",
	Short:   "Start the server",
	PreRunE: config.SetupCobraLogger,
	RunE:    runServer,
}
