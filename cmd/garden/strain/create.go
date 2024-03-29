package strain

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
	createName       string
	createComment    string
	createNotes      string
	createType       string
	createPrice      float64
	createParentage  string
	createTHCPercent float64
	createCBDPercent float64
	createAroma      string
	createTaste      string
	createTags       string
)

func init() {
	createCmd.Flags().StringVar(&createName, "name", "", "Name of a strain")
	createCmd.Flags().StringVar(&createComment, "comment", "", "Comment of a strain")
	createCmd.Flags().StringVar(&createNotes, "notes", "", "Notes of a strain")
	createCmd.Flags().StringVar(&createType, "type", "", "Type of a strain")
	createCmd.Flags().Float64Var(&createPrice, "price", 0, "Price of a strain")
	createCmd.Flags().Float64Var(&createTHCPercent, "thcPercent", 0, "THC percent of a strain")
	createCmd.Flags().Float64Var(&createCBDPercent, "cbdPercent", 0, "CBD percent of a strain")
	createCmd.Flags().StringVar(&createAroma, "aroma", "", "Aroma of a strain")
	createCmd.Flags().StringVar(&createTaste, "taste", "", "Taste of a strain")
	config.Must(createCmd.MarkFlagRequired("name"))
	RootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use: "create",
	Aliases: []string{
		"new",
	},
	Short:   "Create a new strain",
	PreRunE: config.SetupCobraLogger,
	RunE:    createStrain,
}

func createStrain(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	recipeClient := gardenv1connect.NewStrainServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardenv1.CreateStrainRequest{
		Name:       createName,
		Comment:    createComment,
		Notes:      createNotes,
		Type:       createType,
		Price:      createPrice,
		ThcPercent: createTHCPercent,
		CbdPercent: createCBDPercent,
		Parentage:  createParentage,
		Aroma:      createAroma,
		Taste:      createTaste,
		Tags:       createTags,
	})
	resp, err := recipeClient.CreateStrain(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to add strain: %w", err)
	}
	zap.L().Info("Created strain", zap.Any("strain", resp.Msg.Strain))
	return nil
}
