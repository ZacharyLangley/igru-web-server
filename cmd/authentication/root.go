package authentication

import (
	"github.com/ZacharyLangley/igru-web-server/cmd/authentication/group"
	"github.com/ZacharyLangley/igru-web-server/cmd/authentication/session"
	"github.com/ZacharyLangley/igru-web-server/cmd/authentication/user"
	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "authentication",
	Aliases: []string{
		"auth",
	},
}

func init() {
	serveCmd.MarkFlagRequired("config")
	RootCmd.AddCommand(serveCmd)
	RootCmd.AddCommand(user.RootCmd)
	RootCmd.AddCommand(group.RootCmd)
	RootCmd.AddCommand(session.RootCmd)
}

type Config struct {
	Database config.Database `mapstructure:"database"`
	GRPC     config.GRPC     `mapstructure:"grpc"`
	Metrics  config.Metrics  `mapstructure:"metrics"`
}

var serveCmd = &cobra.Command{
	Use:     "serve",
	Short:   "Start the server",
	PreRunE: config.SetupCobraLogger,
	RunE:    runServer,
}
