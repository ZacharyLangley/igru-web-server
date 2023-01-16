package gardens_test

import (
	"log"
	"net/http"
	"testing"

	gardensv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/gardens/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/gardens/v1/gardensv1connect"
	"github.com/bufbuild/connect-go"
)

func initGardensClient() gardensv1connect.GardensServiceClient {
	return gardensv1connect.NewGardensServiceClient(
		http.DefaultClient,
		"http://localhost:80/",
	)
}

func CreateGarden() (*connect.Response[gardensv1.CreateGardenResponse], error) {
	client := initGardensClient()
	req := connect.NewRequest(&gardensv1.CreateGardenRequest{
		Name:          "Garden A_1",
		Comment:       "Garden Mock Comment",
		Location:      "OUTSIDE",
		GrowType:      "SOILLESS",
		GrowSize:      "{\"value\":\"7.15\",\"metric\":\"sq. ft.\"}",
		GrowStyle:     "HYDROPONIC",
		ContainerSize: "{\"value\":\"7.15\",\"metric\":\"sq. ft.\"}",
		Tags:          "[\"Tag A\", \"Tag B\"]",
	})
	ctx, done := initContext()
	defer done()
	return client.CreateGarden(ctx, req)
}

func TestCreateGarden(t *testing.T) {
	res, err := CreateGarden()
	if err != nil {
		t.Fatalf(`Unable to create garden. Error: %q`, err)
	}
	log.Println(res.Msg)
}

func updateGarden(id string) (*connect.Response[gardensv1.UpdateGardenResponse], error) {
	client := initGardensClient()
	req := connect.NewRequest(&gardensv1.UpdateGardenRequest{
		Id:            id,
		Name:          "Garden A_2",
		Comment:       "Updated Garden Mock Comment",
		Location:      "INSIDE",
		GrowType:      "SOIL",
		GrowSize:      "{\"value\":\"7.15\",\"metric\":\"sq. ft.\"}",
		GrowStyle:     "SOIL",
		ContainerSize: "{\"value\":\"7.15\",\"metric\":\"sq. ft.\"}",
		Tags:          "[\"Tag A\", \"Tag B\"]",
	})
	ctx, done := initContext()
	defer done()
	return client.UpdateGarden(ctx, req)
}

func TestUpdateGarden(t *testing.T) {
	res, err := updateGarden("")
	if err != nil {
		t.Fatalf(`Unable to update garden. Error: %q`, err)
	}
	log.Println(res.Msg)
}

func GetGarden(id string) (*connect.Response[gardensv1.GetGardenResponse], error) {
	client := initGardensClient()
	req := connect.NewRequest(&gardensv1.GetGardenRequest{
		Id: id,
	})
	ctx, done := initContext()
	defer done()
	return client.GetGarden(ctx, req)
}
func TestGetGarden(t *testing.T) {
	res, err := GetGarden("")
	if err != nil {
		t.Fatalf(`Unable to get garden. Error: %q`, err)
	}
	log.Println(res.Msg)
}

func GetGardens() (*connect.Response[gardensv1.GetGardensResponse], error) {
	client := initGardensClient()
	req := connect.NewRequest(&gardensv1.GetGardensRequest{})
	ctx, done := initContext()
	defer done()
	return client.GetGardens(ctx, req)
}
func TestGetGardens(t *testing.T) {
	res, err := GetGardens()
	if err != nil {
		t.Fatalf(`Unable to get garden list. Error: %q`, err)
	}
	log.Println(res.Msg)
}

func DeleteGarden(id string) (*connect.Response[gardensv1.DeleteGardenResponse], error) {
	client := initGardensClient()
	req := connect.NewRequest(&gardensv1.DeleteGardenRequest{
		Id: id,
	})
	ctx, done := initContext()
	defer done()
	return client.DeleteGarden(ctx, req)
}
func TestDeleteGarden(t *testing.T) {
	res, err := DeleteGarden("")
	if err != nil {
		t.Fatalf(`Unable to delete garden. Error: %q`, err)
	}
	log.Println(res.Msg)
}
