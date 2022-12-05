package garden

import (
	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "garden",
	Aliases: []string{
		"gardens",
		"g",
	},
	Short: "Utilities for gardens",
}

type Config struct {
	GRPC config.GRPC `mapstructure:"grpc"`
}
