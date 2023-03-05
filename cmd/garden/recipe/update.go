package recipe

import (
	"fmt"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	gardenv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/garden/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/garden/v1/gardenv1connect"
	"github.com/bufbuild/connect-go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	updateID                  string
	updateName                string
	updateComment             string
	updateIngredients         string
	updateInstructions        string
	updatePh                  float64
	updateMixTimeMilliseconds float64
	updateTags                string
)

func init() {
	updateCmd.Flags().StringVar(&updateID, "id", "", "ID of an existing recipe")
	updateCmd.Flags().StringVar(&updateName, "name", "", "Name of a recipe")
	updateCmd.Flags().StringVar(&updateComment, "comment", "", "Comment of a recipe")
	updateCmd.Flags().StringVar(&updateIngredients, "ingredients", "", "Ingredients of a recipe")
	updateCmd.Flags().StringVar(&updateInstructions, "instructions", "", "Instructions of a recipe")
	updateCmd.Flags().Float64Var(&updatePh, "ph", 0, "pH of a recipe")
	updateCmd.Flags().Float64Var(&updateMixTimeMilliseconds, "mixTimeMilliseconds", 0, "MixTimeMilliseconds of a recipe")
	updateCmd.Flags().StringVar(&updateTags, "tags", "", "Tags of a recipe")
	config.Must(updateCmd.MarkFlagRequired("id"))
	RootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update an existing recipe",
	PreRunE: config.SetupCobraLogger,
	RunE:    updateRecipe,
}

func updateRecipe(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	recipeClient := gardenv1connect.NewRecipeServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardenv1.UpdateRecipeRequest{
		Id:                  updateID,
		Name:                updateName,
		Comment:             updateComment,
		Ingredients:         updateIngredients,
		Instructions:        updateInstructions,
		Ph:                  updatePh,
		MixTimeMilliseconds: updateMixTimeMilliseconds,
		Tags:                updateTags,
	})
	resp, err := recipeClient.UpdateRecipe(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to update recipe: %w", err)
	}
	if resp.Msg.Recipe != nil {
		zap.L().Info("updated recipe", zap.Any("recipe", resp.Msg.Recipe))
	}
	return nil
}
