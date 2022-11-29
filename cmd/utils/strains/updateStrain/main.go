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
	client := gardensv1connect.NewStrainsServiceClient(
		http.DefaultClient,
		"http://localhost:8082/",
	)
	req := connect.NewRequest(&gardensv1.UpdateStrainRequest{
		Id:         "e618113a-f71a-484f-b694-f37cc2dc97d4",
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
	ctx, done := context.WithTimeout(context.Background(), time.Second*5)
	defer done()
	res, err := client.UpdateStrain(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}
