package main

import (
	"fmt"
	"net/http"

	"github.com/danielbintar/angel/server/users"
	"github.com/danielbintar/angel/server/users/db"
	"github.com/danielbintar/angel/server/users/router"

	"github.com/danielbintar/angel/server-library/pubsub"
)

func main() {
	r := router.NewRouter()

	database := db.NewDB()
	publisher := pubsub.NewKafkaAsyncProducer()
	m := users.UserManager {
		DatabaseManager: database,
		Publisher: publisher,
	}

	router.Public(r, &m)

	fmt.Println("listen to 7001")
	http.ListenAndServe(":7001", r)
}
