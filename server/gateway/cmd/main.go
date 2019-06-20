package main

import (
	"log"
	"net/http"
	"time"

	"github.com/danielbintar/angel/server/gateway"
)

func main() {
	s := gateway.NewServer()

	h := &http.Server{
		Addr:         ":8089",
		Handler:      s,
		ReadTimeout:  310 * time.Second,
		WriteTimeout: 310 * time.Second,
	}

	log.Println("Listening to 8089")
	if err := h.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
