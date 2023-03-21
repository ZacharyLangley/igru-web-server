package ingress

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "ingress",
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
