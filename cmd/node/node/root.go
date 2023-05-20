package node

import (
	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "node",
	Aliases: []string{
		"nodes",
		"n",
	},
	Short: "Utilities for nodes",
}

type Config struct {
	GRPC config.GRPC `mapstructure:"grpc"`
}
