package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
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
	req := connect.NewRequest(&authenticationv1.WhoamiRequest{})
	authentication.AddSessionToken(req.Header(), os.Args[1])
	ctx, done := context.WithTimeout(context.Background(), time.Second*5)
	defer done()
	res, err := client.Whoami(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.Msg)
}
