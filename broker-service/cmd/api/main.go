package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Println("Broker Server Started on port:", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routers(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
