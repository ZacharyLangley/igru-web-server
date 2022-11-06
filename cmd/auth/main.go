package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/connect"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/authentication"
)

const webPort = "80"

var counts int64

func main() {
	log.Println("Starting authentication service")

	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	mux := connect.CreateMux(authentication.New(conn))
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: mux,
	}

	// Start cleanup process
	go func(ctx context.Context) {
		t := time.NewTicker(time.Second * 10)
		for range t.C {
			queries := models.New(conn)
			sessions, err := queries.GetExpiredSessions(context.Background())
			if err != nil {
				log.Println("Failed to get expired sessions", err)
			}
			log.Println("Clearing old sessions", len(sessions))
			for _, session := range sessions {
				if session.ExpiredAt.Before(time.Now()) {
					log.Println("Deleting", session.ID)
					if err := queries.DeleteSession(context.Background(), session.ID); err != nil {
						log.Println("Failed to delete expired sessions")
					}
				}
			}
		}
	}(context.Background())

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := database.Open(context.Background(), dsn)
		if err != nil {
			log.Println("Postgres not yet ready ...")
			log.Println(err)
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}
