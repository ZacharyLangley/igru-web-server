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
)

var (
	deleteRecipeID string
)

func init() {
	deleteCmd.Flags().StringVar(&deleteRecipeID, "id", "", "ID of an existing recipe")
	config.Must(deleteCmd.MarkFlagRequired("id"))
	RootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use: "delete",
	Aliases: []string{
		"rm",
		"del",
	},
	Short:   "Delete an existing recipe",
	PreRunE: config.SetupCobraLogger,
	RunE:    deleteRecipe,
}

func deleteRecipe(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	recipeClient := gardenv1connect.NewRecipeServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardenv1.DeleteRecipeRequest{
		Id: deleteRecipeID,
	})
	_, err := recipeClient.DeleteRecipe(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to delete recipe: %w", err)
	}
	return nil
}
