package cmd

import (
	"os"

	"github.com/ZacharyLangley/igru-web-server/cmd/authentication"
	"github.com/ZacharyLangley/igru-web-server/cmd/garden"
	"github.com/ZacharyLangley/igru-web-server/cmd/ingress"
	"github.com/ZacharyLangley/igru-web-server/cmd/node"
	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "igru-web-server",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Bool("dev-logs", false, "Development formatted logs")
	rootCmd.PersistentFlags().String("level", "info", "Level of verbosity to log")
	rootCmd.PersistentFlags().String("config", "", "config file")
	config.Must(viper.BindPFlags(rootCmd.PersistentFlags()))

	rootCmd.AddCommand(authentication.RootCmd)
	rootCmd.AddCommand(garden.RootCmd)
	rootCmd.AddCommand(ingress.RootCmd)
	rootCmd.AddCommand(node.RootCmd)
	rootCmd.AddCommand(healthCmd)
}
