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
	createName            string
	createComment         string
	createNotes           string
	createGrowCycleLength string
	createParentage       string
	createOrigin          string
	createGender          string
	createDaysFlowering   float64
	createDaysCured       float64
	createHarvestWeight   string
	createBudDensity      float64
	createBudPistils      bool
	createTags            string
	createSetAcquiredAt   bool
)

func init() {
	createCmd.Flags().StringVar(&createName, "name", "", "Name of a new plant")
	createCmd.Flags().StringVar(&createComment, "comment", "", "Comment of a new plant")
	createCmd.Flags().StringVar(&createNotes, "notes", "", "Notes of a new plant")
	createCmd.Flags().StringVar(&createGrowCycleLength, "growCyleLength", "", "Grow cycle length of a new plant")
	createCmd.Flags().StringVar(&createParentage, "parentage", "", "Parentage of a new plant")
	createCmd.Flags().StringVar(&createOrigin, "origin", "", "Origin of a new plant")
	createCmd.Flags().StringVar(&createGender, "gender", "", "Gender of a new plant")
	createCmd.Flags().Float64Var(&createDaysFlowering, "daysFlowering", 0, "Days flowering of a new plant")
	createCmd.Flags().Float64Var(&createDaysCured, "daysCured", 0, "Days cured of a new plant")
	createCmd.Flags().StringVar(&createHarvestWeight, "harvestWeight", "", "Harvest weight of a new plant")
	createCmd.Flags().Float64Var(&createBudDensity, "budDensity", 0, "Bud density of a new plant")
	createCmd.Flags().BoolVar(&createBudPistils, "budPistils", false, "Bud pistils of a new plant")
	createCmd.Flags().StringVar(&createTags, "tags", "", "Tags of a new plant")
	createCmd.Flags().BoolVar(&createSetAcquiredAt, "setAcquiredAt", false, "Set acquired at time of an existing plant to now")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("config")
	RootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use: "create",
	Aliases: []string{
		"new",
	},
	Short:   "Create a new plant",
	PreRunE: config.SetupCobraLogger,
	RunE:    createGarden,
}

func createGarden(cmd *cobra.Command, args []string) error {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	userClient := gardensv1connect.NewPlantsServiceClient(
		http.DefaultClient,
		cfg.GRPC.Address,
	)
	ctx := context.New(cmd.Context())
	req := connect.NewRequest(&gardensv1.CreatePlantRequest{
		Name:            createName,
		Comment:         createComment,
		Notes:           createNotes,
		GrowCycleLength: createGrowCycleLength,
		Parentage:       createParentage,
		Origin:          createOrigin,
		Gender:          createGender,
		DaysFlowering:   createDaysFlowering,
		DaysCured:       createDaysCured,
		HarvestedWeight: createHarvestWeight,
		BudDensity:      createBudDensity,
		BudPistils:      createBudPistils,
		Tags:            createTags,
	})
	if createSetAcquiredAt {
		req.Msg.AcquiredAt = timestamppb.Now()
	}
	resp, err := userClient.CreatePlant(ctx, req)
	if err != nil {
		return fmt.Errorf("Failed to add plant: %w", err)
	}
	zap.L().Info("Created plant", zap.Any("plant", resp.Msg.Plant))
	return nil
}
