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
	req := connect.NewRequest(&gardensv1.UpdateGardenRequest{
		Id:            "c299e110-3944-4b4e-a5a7-af8cab0e8bc7",
		Name:          "Garden A_2",
		Comment:       "Updated Garden Mock Comment",
		Location:      "INSIDE",
		GrowType:      "SOIL",
		GrowSize:      "{\"value\":\"7.15\",\"metric\":\"sq. ft.\"}",
		GrowStyle:     "SOIL",
		ContainerSize: "{\"value\":\"7.15\",\"metric\":\"sq. ft.\"}",
		Tags:          "[\"Tag A\", \"Tag B\"]",
	})
	ctx, done := context.WithTimeout(context.Background(), time.Second*5)
	defer done()
	res, err := client.UpdateGarden(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}
