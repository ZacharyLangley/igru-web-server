package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/broker"
)

const webPort = "80"

type Config struct{}

func main() {
	log.Println("Broker Server Started on port:", webPort)

	mux := connect.CreateMux(broker.New())
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
