package garden

import (
	"github.com/ZacharyLangley/igru-web-server/cmd/garden/garden"
	"github.com/ZacharyLangley/igru-web-server/cmd/garden/plant"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "garden",
	Aliases: []string{
		"gardens",
	},
}

func init() {
	serveCmd.MarkFlagRequired("config")
	RootCmd.AddCommand(serveCmd)
	RootCmd.AddCommand(garden.RootCmd)
	RootCmd.AddCommand(plant.RootCmd)
}
