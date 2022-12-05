package plant

import (
	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "plant",
	Aliases: []string{
		"plants",
		"p",
	},
	Short: "Utilities for plants",
}

type Config struct {
	GRPC config.GRPC `mapstructure:"grpc"`
}
