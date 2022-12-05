package member

import (
	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "member",
	Aliases: []string{
		"members",
		"m",
	},
	Short: "Utilities for user group members",
}

type Config struct {
	GRPC config.GRPC `mapstructure:"grpc"`
}
