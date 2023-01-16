package session

import (
	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "session",
	Aliases: []string{
		"sessions",
		"sess",
		"s",
	},
	Short: "Utilities for sessions",
}

type Config struct {
	GRPC config.GRPC `mapstructure:"grpc"`
}
