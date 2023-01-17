package gardens_test

import (
	"log"
	"net/http"
	"testing"

	gardenv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/garden/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/garden/v1/gardenv1connect"
	"github.com/bufbuild/connect-go"
)

func initPlantsClient() gardenv1connect.PlantServiceClient {
	return gardenv1connect.NewPlantServiceClient(
		http.DefaultClient,
		"http://localhost:80/",
	)
}

func CreatePlant() (*connect.Response[gardenv1.CreatePlantResponse], error) {
	client := initPlantsClient()
	req := connect.NewRequest(&gardenv1.CreatePlantRequest{
		Name:            "Plant A_1",
		Comment:         "Plant Mock Comment",
		Notes:           "Plant Mock Note",
		GrowCycleLength: "{\"value\":\"28\",\"metric\":\"days\"}",
		Parentage:       "Mock Parent Strain",
		Origin:          "Clone",
		Gender:          "Feminized",
		DaysFlowering:   2.4,
		DaysCured:       1.2,
		HarvestedWeight: "{\"value\":\"1.05\",\"metric\":\"lbs.\"}",
		BudDensity:      0.7,
		BudPistils:      false,
		Tags:            "[\"Tag A\", \"Tag B\"]",
	})
	ctx, done := initContext()
	defer done()
	return client.CreatePlant(ctx, req)
}

func TestCreatePlant(t *testing.T) {
	res, err := CreatePlant()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}

func UpdatePlant(id string) (*connect.Response[gardenv1.UpdatePlantResponse], error) {
	client := initPlantsClient()
	req := connect.NewRequest(&gardenv1.UpdatePlantRequest{
		Id:              id,
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
	ctx, done := initContext()
	defer done()
	return client.UpdatePlant(ctx, req)
}

func TestUpdatePlant(t *testing.T) {
	res, err := UpdatePlant("")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}

func GetPlant(id string) (*connect.Response[gardenv1.GetPlantResponse], error) {
	client := initPlantsClient()
	req := connect.NewRequest(&gardenv1.GetPlantRequest{
		Id: id,
	})
	ctx, done := initContext()
	defer done()
	return client.GetPlant(ctx, req)
}

func TestGetPlant(t *testing.T) {
	res, err := GetPlant("")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}

func GetPlants() (*connect.Response[gardenv1.GetPlantsResponse], error) {
	client := initPlantsClient()
	req := connect.NewRequest(&gardenv1.GetPlantsRequest{})
	ctx, done := initContext()
	defer done()
	return client.GetPlants(ctx, req)
}

func TestGetPlants(t *testing.T) {
	res, err := GetPlants()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}

func DeletePlant(id string) (*connect.Response[gardenv1.DeletePlantResponse], error) {
	client := initPlantsClient()
	req := connect.NewRequest(&gardenv1.DeletePlantRequest{
		Id: "INSERT-NEW-UUID-HERE",
	})
	ctx, done := initContext()
	defer done()
	return client.DeletePlant(ctx, req)
}

func TestDeletePlant(t *testing.T) {
	res, err := DeletePlant("")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}
