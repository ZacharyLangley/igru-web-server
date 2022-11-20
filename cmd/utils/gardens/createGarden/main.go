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
	client := gardensv1connect.NewGardensServiceClient(
		http.DefaultClient,
		"http://localhost:8082/",
	)
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
	ctx, done := context.WithTimeout(context.Background(), time.Second*5)
	defer done()
	res, err := client.CreateGarden(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}
