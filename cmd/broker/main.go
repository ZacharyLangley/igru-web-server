package main

import (
	"log"
	"net/http"

	"github.com/ZacharyLangley/igru-web-server/pkg/config"
	"github.com/ZacharyLangley/igru-web-server/pkg/connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/broker"
)

type Config struct {
	GRPC config.GRPC `mapstructure:"grpc"`
}

func main() {
	var cfg Config
	if err := config.New(&cfg); err != nil {
		log.Fatalln(err)
	}
	log.Println("Broker Server Started")

	mux := connect.CreateMux(broker.New())
	srv := &http.Server{
		Addr:    cfg.GRPC.Address,
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
