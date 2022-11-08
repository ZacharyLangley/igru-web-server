package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	authenticationv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/authentication"
	"github.com/bufbuild/connect-go"
)

func main() {
	client := authenticationv1connect.NewUserServiceClient(
		http.DefaultClient,
		"http://localhost:8081/",
	)
	req := connect.NewRequest(&authenticationv1.AuthenticateRequest{
		Email:    "bobcob333@hotmail.com",
		Password: "password123",
	})
	ctx, done := context.WithTimeout(context.Background(), time.Second*5)
	defer done()
	res, err := client.Authenticate(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	sessionID, err := authentication.ExtractSessionToken(res.Header())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(sessionID)
}
