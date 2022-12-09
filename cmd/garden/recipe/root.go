package recipe

import (
	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "recipe",
	Aliases: []string{
		"recipes",
		"r",
	},
	Short: "Utilities for recipes",
}

type Config struct {
	GRPC config.GRPC `mapstructure:"grpc"`
}
