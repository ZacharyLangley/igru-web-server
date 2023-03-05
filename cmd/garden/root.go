package garden

import (
	"github.com/ZacharyLangley/igru-web-server/cmd/garden/garden"
	"github.com/ZacharyLangley/igru-web-server/cmd/garden/plant"
	"github.com/ZacharyLangley/igru-web-server/cmd/garden/recipe"
	"github.com/ZacharyLangley/igru-web-server/cmd/garden/strain"
	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "garden",
	Aliases: []string{
		"gardens",
	},
}

func init() {
	config.Must(serveCmd.MarkFlagRequired("config"))
	RootCmd.AddCommand(serveCmd)
	RootCmd.AddCommand(garden.RootCmd)
	RootCmd.AddCommand(plant.RootCmd)
	RootCmd.AddCommand(strain.RootCmd)
	RootCmd.AddCommand(recipe.RootCmd)
}
