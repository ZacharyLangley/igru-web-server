package user

import (
	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "user",
	Aliases: []string{
		"users",
		"usr",
		"u",
	},
	Short: "Utilities for users",
}

type Config struct {
	GRPC config.GRPC `mapstructure:"grpc"`
}
