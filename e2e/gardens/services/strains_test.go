package gardens_test

import (
	"log"
	"net/http"
	"testing"

	gardensv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/gardens/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/gardens/v1/gardensv1connect"
	"github.com/bufbuild/connect-go"
)

func initStrainsClient() gardensv1connect.StrainsServiceClient {
	return gardensv1connect.NewStrainsServiceClient(
		http.DefaultClient,
		"http://localhost:80/",
	)
}

func CreateStrain() (*connect.Response[gardensv1.CreateStrainResponse], error) {
	client := initStrainsClient()
	req := connect.NewRequest(&gardensv1.CreateStrainRequest{
		Name:       "Strain A_1",
		Comment:    "Strain Mock Comment",
		Notes:      "Strain Mock Note",
		Type:       "SATIVA_HYBRID",
		Price:      4.08,
		ThcPercent: 0.68,
		CbdPercent: 0.32,
		Parentage:  "PARENT-STRAIN-UUID",
		Aroma:      "[\"Light\", \"Citrus\"]",
		Taste:      "[\"Citrus\", \"Hard\"]",
		Tags:       "[\"Tag A\", \"Tag B\"]",
	})
	ctx, done := initContext()
	defer done()
	return client.CreateStrain(ctx, req)
}

func TestCreateStrain(t *testing.T) {
	res, err := CreateStrain()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}

func UpdateStrain(id string) (*connect.Response[gardensv1.UpdateStrainResponse], error) {
	client := initStrainsClient()
	req := connect.NewRequest(&gardensv1.UpdateStrainRequest{
		Id:         id,
		Name:       "Updated Strain A_1",
		Comment:    "Updated Strain Mock Comment",
		Notes:      "Updated Strain Mock Note",
		Type:       "SATIVA",
		Price:      5.18,
		ThcPercent: 0.50,
		CbdPercent: 0.50,
		Parentage:  "UPDATED-PARENT-STRAIN-UUID",
		Aroma:      "[\"Updated\", \"Light\", \"Orange\"]",
		Taste:      "[\"Updated\", \"Citrus\", \"Hard\"]",
		Tags:       "[\"Updated Tag A\", \"Updated Tag B\"]",
	})
	ctx, done := initContext()
	defer done()
	return client.UpdateStrain(ctx, req)
}

func TestUpdateStrain(t *testing.T) {
	res, err := UpdateStrain("")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}

func GetStrain(id string) (*connect.Response[gardensv1.GetStrainResponse], error) {
	client := initStrainsClient()
	req := connect.NewRequest(&gardensv1.GetStrainRequest{
		Id: id,
	})
	ctx, done := initContext()
	defer done()
	return client.GetStrain(ctx, req)
}

func TestGetStrain(t *testing.T) {
	res, err := GetStrain("")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}

func GetStrains() (*connect.Response[gardensv1.GetStrainsResponse], error) {
	client := initStrainsClient()
	req := connect.NewRequest(&gardensv1.GetStrainsRequest{})
	ctx, done := initContext()
	defer done()
	return client.GetStrains(ctx, req)
}

func TestGetStrains(t *testing.T) {
	res, err := GetStrains()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}

func DeleteStrain(id string) (*connect.Response[gardensv1.DeleteStrainResponse], error) {
	client := initStrainsClient()
	req := connect.NewRequest(&gardensv1.DeleteStrainRequest{
		Id: id,
	})
	ctx, done := initContext()
	defer done()
	return client.DeleteStrain(ctx, req)
}

func TestDeleteStrain(t *testing.T) {
	res, err := DeleteStrain("")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}
