package ingress

import (
	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "ingress",
}

func init() {
	config.Must(serveCmd.MarkFlagRequired("config"))
	RootCmd.AddCommand(serveCmd)
}
