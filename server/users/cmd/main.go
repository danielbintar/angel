package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/danielbintar/angel/server/users"
	"github.com/danielbintar/angel/server/users/db"
	"github.com/danielbintar/angel/server/users/router"

	"github.com/danielbintar/angel/server-library/pubsub"
)

func main() {
	defer func() {
		log.Println("Server exited properly")
	}()

	r := router.NewRouter()

	database := db.NewDB()
	defer func() {
		log.Println("Closing database")
		err := database.Close()
		if err != nil {
			log.Println("Failed: " + err.Error())
		}
	}()

	publisher := pubsub.NewKafkaAsyncProducer()
	defer func() {
		log.Println("Closing kafka publisher")
		err := publisher.Close()
		if err != nil {
			log.Println("Failed: " + err.Error())
		}
	}()

	m := users.UserManager {
		DatabaseManager: database,
		Publisher: publisher,
	}

	router.Public(r, &m)

	srv := &http.Server{
		Addr:    ":7001",
		Handler: r,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("listen failed: " + err.Error())
			return
		}
	}()

	log.Println("server users listen to 7001")

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 16 * time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server shutdown failed: ", err.Error())
	}
}
