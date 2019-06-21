package main

import (
	"os"

	"github.com/danielbintar/angel/server-library/migration"
)

func main() {
	key := "POSTGRES_DATABASE"
	if os.Getenv("ENVIRONMENT") == "test" {
		key = "TEST_" + key
	}

	dbName := os.Getenv(key)

	var query string
	if len(os.Args) == 2 && os.Args[1] == "down" {
		query = "DROP DATABASE IF EXISTS " + dbName
	} else {
		query = "CREATE DATABASE " + dbName
	}

	migration.RunPostgres(&migration.PostgresQueryOpt{Query: query, Base: true})
}
