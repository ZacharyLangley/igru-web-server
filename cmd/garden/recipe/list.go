package recipe

import (
	"fmt"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	gardensv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/gardens/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/gardens/v1/gardensv1connect"
	"github.com/bufbuild/connect-go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use: "list",
	Aliases: []string{
		"l",
	},
	Short:   "Get all existing recipes",
	PreRunE: config.SetupCobraLogger,
	RunE:    getRecipes,
}

func getRecipes(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	recipeClient := gardensv1connect.NewRecipesServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardensv1.GetRecipesRequest{})
	resp, err := recipeClient.GetRecipes(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get recipes: %w", err)
	}
	if resp.Msg.Recipes != nil {
		zap.L().Info("Found recipes", zap.Any("recipes", resp.Msg.Recipes))
	}
	return nil
}
