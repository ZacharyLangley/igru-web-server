package main

import (
	"context"
	"log"
	"net/http"
	"time"

	authenticationv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/bufbuild/connect-go"
)

func main() {
	client := authenticationv1connect.NewUserServiceClient(
		http.DefaultClient,
		"http://localhost:8081/",
	)
	req := connect.NewRequest(&authenticationv1.CreateRequest{
		FirstName: "Eric",
		LastName:  "Suedmeier",
		Email:     "bobcob7@hotmail.com",
		Password:  "password123",
	})
	ctx, done := context.WithTimeout(context.Background(), time.Second*5)
	defer done()
	res, err := client.Create(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Msg)
}
