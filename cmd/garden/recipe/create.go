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

var (
	createName                string
	createComment             string
	createIngredients         string
	createInstructions        string
	createPh                  float64
	createMixTimeMilliseconds float64
	createTags                string
)

func init() {
	createCmd.Flags().StringVar(&createName, "name", "", "Name of a recipe")
	createCmd.Flags().StringVar(&createComment, "comment", "", "Comment of a recipe")
	createCmd.Flags().StringVar(&createIngredients, "ingredients", "", "Ingredients of a recipe")
	createCmd.Flags().StringVar(&createInstructions, "instructions", "", "Instructions of a recipe")
	createCmd.Flags().Float64Var(&createPh, "ph", 0, "pH of a recipe")
	createCmd.Flags().Float64Var(&createMixTimeMilliseconds, "mixTimeMilliseconds", 0, "MixTimeMilliseconds of a recipe")
	createCmd.Flags().StringVar(&createTags, "tags", "", "Tags of a recipe")
	createCmd.MarkFlagRequired("name")
	RootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use: "create",
	Aliases: []string{
		"new",
	},
	Short:   "Create a new recipe",
	PreRunE: config.SetupCobraLogger,
	RunE:    createRecipe,
}

func createRecipe(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	recipeClient := gardensv1connect.NewRecipesServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardensv1.CreateRecipeRequest{
		Name:                createName,
		Comment:             createComment,
		Ingredients:         createIngredients,
		Instructions:        createInstructions,
		Ph:                  createPh,
		MixTimeMilliseconds: createMixTimeMilliseconds,
		Tags:                createTags,
	})
	resp, err := recipeClient.CreateRecipe(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to add recipe: %w", err)
	}
	zap.L().Info("Created recipe", zap.Any("recipe", resp.Msg.Recipe))
	return nil
}
