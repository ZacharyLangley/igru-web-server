package ingress

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "ingress",
}

func init() {
	serveCmd.MarkFlagRequired("config")
	RootCmd.AddCommand(serveCmd)
}
