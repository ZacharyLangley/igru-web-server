package group

import (
	"github.com/ZacharyLangley/igru-web-server/cmd/authentication/group/member"
	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "group",
	Aliases: []string{
		"groups",
		"grp",
		"g",
	},
	Short: "Utilities for user groups",
}

type Config struct {
	GRPC config.GRPC `mapstructure:"grpc"`
}

func init() {
	RootCmd.AddCommand(member.RootCmd)
}
