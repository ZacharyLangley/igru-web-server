package broker

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "broker",
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
