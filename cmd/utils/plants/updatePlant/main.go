package main

import (
	"context"
	"log"
	"net/http"
	"time"

	gardensv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/gardens/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/gardens/v1/gardensv1connect"
	"github.com/bufbuild/connect-go"
)

func main() {
	client := gardensv1connect.NewPlantsServiceClient(
		http.DefaultClient,
		"http://localhost:8082/",
	)
	req := connect.NewRequest(&gardensv1.UpdatePlantRequest{
		Id:              "INSERT-NEW-UUID-HERE",
		Name:            "Plant A_2",
		Comment:         "Plant Mock Updated Comment",
		Notes:           "Plant Mock Updated Note",
		GrowCycleLength: "{\"value\":\"28\",\"metric\":\"days\"}",
		Parentage:       "Mock Parent Strain",
		Origin:          "Clone",
		Gender:          "Feminized",
		DaysFlowering:   22.4,
		DaysCured:       0,
		HarvestedWeight: "{\"value\":\"1.05\",\"metric\":\"lbs.\"}",
		BudDensity:      0.7,
		BudPistils:      true,
		Tags:            "[\"Tag A\", \"Tag B\"]",
	})
	ctx, done := context.WithTimeout(context.Background(), time.Second*5)
	defer done()
	res, err := client.UpdatePlant(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}
