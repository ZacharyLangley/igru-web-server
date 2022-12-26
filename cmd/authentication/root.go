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
	createKeyCmd.Flags().StringVar(&createKeyFile, "file", "", "Full name of a new key")
	createKeyCmd.Flags().StringVar(&createKeyType, "type", "ed25519", "Email of a new key")
	createKeyCmd.MarkFlagRequired("file")
	createKeySetCmd.Flags().StringVar(&createKeySetFile, "file", "", "Full name of a new key")
	createKeySetCmd.Flags().StringVar(&createKeySetType, "type", "ed25519", "Email of a new key")
	createKeySetCmd.MarkFlagRequired("file")
	RootCmd.AddCommand(createKeyCmd)
	RootCmd.AddCommand(createKeySetCmd)
	RootCmd.AddCommand(user.RootCmd)
	RootCmd.AddCommand(group.RootCmd)
	RootCmd.AddCommand(session.RootCmd)
}

type Config struct {
	Database   config.Database   `mapstructure:"database"`
	GRPC       config.GRPC       `mapstructure:"grpc"`
	Metrics    config.Metrics    `mapstructure:"metrics"`
	AuthServer config.AuthServer `mapstructure:"authServer"`
}

var serveCmd = &cobra.Command{
	Use:     "serve",
	Short:   "Start the server",
	PreRunE: config.SetupCobraLogger,
	RunE:    runServer,
}

var createKeyCmd = &cobra.Command{
	Use:     "create-key",
	Short:   "Create new key pair for server",
	PreRunE: config.SetupCobraLogger,
	RunE:    createKey,
}

var createKeySetCmd = &cobra.Command{
	Use:     "create-key-set",
	Short:   "Create new key pair set for server",
	PreRunE: config.SetupCobraLogger,
	RunE:    createKeySet,
}
