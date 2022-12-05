package strain

import (
	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "strain",
	Aliases: []string{
		"strains",
		"s",
	},
	Short: "Utilities for strains",
}

type Config struct {
	GRPC config.GRPC `mapstructure:"grpc"`
}
