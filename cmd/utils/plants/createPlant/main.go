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
	req := connect.NewRequest(&gardensv1.CreatePlantRequest{
		Name:            "Plant A_1",
		Comment:         "Plant Mock Comment",
		Notes:           "Plant Mock Note",
		GrowCycleLength: "{\"value\":\"28\",\"metric\":\"days\"}",
		Parentage:       "INSERT-PARENTAGE-PLANT-ID",
		Origin:          "Clone",
		Gender:          "Feminized",
		DaysFlowering:   2.4,
		DaysCured:       1.2,
		HarvestedWeight: "{\"value\":\"1.05\",\"metric\":\"lbs.\"}",
		BudDensity:      0.7,
		BudPistils:      false,
		Tags:            "[\"Tag A\", \"Tag B\"]",
	})
	ctx, done := context.WithTimeout(context.Background(), time.Second*5)
	defer done()
	res, err := client.CreatePlant(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}
