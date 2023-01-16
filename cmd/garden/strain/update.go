package strain

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
	updateID         string
	updateName       string
	updateComment    string
	updateNotes      string
	updateType       string
	updatePrice      float64
	updateParentage  string
	updateTHCPercent float64
	updateCBDPercent float64
	updateAroma      string
	updateTaste      string
	updateTags       string
)

func init() {
	updateCmd.Flags().StringVar(&updateID, "id", "", "ID of an existing strain")
	updateCmd.Flags().StringVar(&updateName, "name", "", "Name of an existing strain")
	updateCmd.Flags().StringVar(&updateComment, "comment", "", "Comment of an existing strain")
	updateCmd.Flags().StringVar(&updateNotes, "notes", "", "Notes of an existing strain")
	updateCmd.Flags().StringVar(&updateType, "type", "", "Type of an existing strain")
	updateCmd.Flags().Float64Var(&updatePrice, "price", 0, "Price of an existing strain")
	updateCmd.Flags().Float64Var(&updateTHCPercent, "thcPercent", 0, "THC percent of an existing strain")
	updateCmd.Flags().Float64Var(&updateCBDPercent, "cbdPercent", 0, "CBD percent of an existing strain")
	updateCmd.Flags().StringVar(&updateAroma, "aroma", "", "Aroma of an existing strain")
	updateCmd.Flags().StringVar(&updateTaste, "taste", "", "Taste of an existing strain")
	updateCmd.Flags().StringVar(&updateTags, "tags", "", "Tags of an existing strain")
	updateCmd.MarkFlagRequired("id")
	updateCmd.MarkFlagRequired("config")
	RootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update an existing strain",
	PreRunE: config.SetupCobraLogger,
	RunE:    updateStrain,
}

func updateStrain(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	recipeClient := gardensv1connect.NewStrainsServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardensv1.UpdateStrainRequest{
		Id:         updateID,
		Name:       updateName,
		Comment:    updateComment,
		Notes:      updateNotes,
		Type:       updateType,
		Price:      updatePrice,
		ThcPercent: updateTHCPercent,
		CbdPercent: updateCBDPercent,
		Parentage:  updateParentage,
		Aroma:      updateAroma,
		Taste:      updateTaste,
		Tags:       updateTags,
	})
	resp, err := recipeClient.UpdateStrain(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to update strain: %w", err)
	}
	if resp.Msg.Strain != nil {
		zap.L().Info("updated strain", zap.Any("strian", resp.Msg.Strain))
	}
	return nil
}
