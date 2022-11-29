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
	ctx, done := context.WithTimeout(context.Background(), time.Second*5)
	defer done()
	res, err := client.CreateStrain(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}
