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
	getRecipeID string
)

func init() {
	getCmd.Flags().StringVar(&getRecipeID, "id", "", "id of an existing recipe")
	getCmd.MarkFlagRequired("id")
	RootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use: "get",
	Aliases: []string{
		"g",
	},
	Short:   "Get an existing recipe",
	PreRunE: config.SetupCobraLogger,
	RunE:    getRecipe,
}

func getRecipe(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	recipeClient := gardenv1connect.NewRecipeServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardenv1.GetRecipeRequest{
		Id: getRecipeID,
	})
	resp, err := recipeClient.GetRecipe(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to get recipe: %w", err)
	}
	if resp.Msg.Recipe != nil {
		zap.L().Info("Found recipe", zap.Any("recipe", resp.Msg.Recipe))
	}
	return nil
}
