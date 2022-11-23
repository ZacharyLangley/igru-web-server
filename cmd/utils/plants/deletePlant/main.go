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
	req := connect.NewRequest(&gardensv1.DeletePlantRequest{
		Id: "INSERT-NEW-UUID-HERE",
	})
	ctx, done := context.WithTimeout(context.Background(), time.Second*5)
	defer done()
	res, err := client.DeletePlant(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}
