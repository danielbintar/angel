package main

import (
	"os"

	"github.com/danielbintar/angel/server/users/db/migration"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	key := "MYSQL_DATABASE"
	if os.Getenv("ENVIRONMENT") == "TEST" { key = "TEST_" + key }
	dbName := os.Getenv(key)

	var query string
	if len(os.Args) == 2 && os.Args[1] == "down" {
		query = "DROP DATABASE IF EXISTS "+ dbName
	} else {
		query = "CREATE DATABASE IF NOT EXISTS "+ dbName
	}

	migration.Run(&migration.QueryOpt { Query: query, Base: true })
}
