package main

import (
	"fmt"
	"net/http"

	"github.com/danielbintar/angel/server/users"
	"github.com/danielbintar/angel/server/users/db"
	"github.com/danielbintar/angel/server/users/router"
)

func main() {
	r := router.NewRouter()

	database := db.NewDB()
	m := users.UserManager { DatabaseManager: database }

	router.Public(r, &m)

	fmt.Println("listen to 7001")
	http.ListenAndServe(":7001", r)
}
