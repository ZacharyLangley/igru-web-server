package plant

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
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	updateID              string
	updateName            string
	updateComment         string
	updateNotes           string
	updateGrowCycleLength string
	updateParentage       string
	updateOrigin          string
	updateGender          string
	updateDaysFlowering   float64
	updateDaysCured       float64
	updateHarvestWeight   string
	updateBudDensity      float64
	updateBudPistils      bool
	updateTags            string
	updateSetAcquiredAt   bool
)

func init() {
	updateCmd.Flags().StringVar(&updateID, "id", "", "ID of an existing plant")
	updateCmd.Flags().StringVar(&updateName, "name", "", "Name of an existing plant")
	updateCmd.Flags().StringVar(&updateComment, "comment", "", "Comment of an existing plant")
	updateCmd.Flags().StringVar(&updateNotes, "notes", "", "Notes of an existing plant")
	updateCmd.Flags().StringVar(&updateGrowCycleLength, "growCycleLength", "", "Grow cycle length of an existing plant")
	updateCmd.Flags().StringVar(&updateParentage, "parentage", "", "Parentage of an existing plant")
	updateCmd.Flags().StringVar(&updateOrigin, "origin", "", "Origin of an existing plant")
	updateCmd.Flags().StringVar(&updateGender, "gender", "", "Gender of an existing plant")
	updateCmd.Flags().Float64Var(&updateDaysFlowering, "daysFlowering", 0, "Days flowering of an existing plant")
	updateCmd.Flags().Float64Var(&updateDaysCured, "daysCured", 0, "Grow size of an existing plant")
	updateCmd.Flags().StringVar(&updateHarvestWeight, "harvestWeight", "", "Harvested weight of an existing plant")
	updateCmd.Flags().Float64Var(&updateBudDensity, "budDensity", 0, "Bud density of an existing plant")
	updateCmd.Flags().BoolVar(&updateBudPistils, "budPistils", false, "Bud Pistils of an existing plant")
	updateCmd.Flags().StringVar(&updateTags, "tags", "", "Tags of an existing plant")
	updateCmd.Flags().BoolVar(&updateSetAcquiredAt, "setAcquiredAt", false, "Set acquired at time of an existing plant to now")
	updateCmd.MarkFlagRequired("name")
	RootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update an existing plants",
	PreRunE: config.SetupCobraLogger,
	RunE:    updatePlant,
}

func updatePlant(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	plantClient := gardensv1connect.NewPlantsServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardensv1.UpdatePlantRequest{
		Id:              updateID,
		Name:            updateName,
		Comment:         updateComment,
		Notes:           updateNotes,
		GrowCycleLength: updateGrowCycleLength,
		Parentage:       updateParentage,
		Origin:          updateOrigin,
		Gender:          updateGender,
		DaysFlowering:   updateDaysFlowering,
		DaysCured:       updateDaysCured,
		HarvestedWeight: updateHarvestWeight,
		BudDensity:      updateBudDensity,
		BudPistils:      updateBudPistils,
		Tags:            updateTags,
	})
	if updateSetAcquiredAt {
		req.Msg.AcquiredAt = timestamppb.Now()
	}
	resp, err := plantClient.UpdatePlant(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to update plant: %w", err)
	}
	if resp.Msg.Plant != nil {
		zap.L().Info("updated plant", zap.Any("plant", resp.Msg.Plant))
	}
	return nil
}
