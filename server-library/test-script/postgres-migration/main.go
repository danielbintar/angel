package main

import (
	"os"

	"github.com/danielbintar/angel/server-library/migration"
)

func main() {
	key := "TEST_POSTGRES_DATABASE"
	dbName := os.Getenv(key)

	migration.RunPostgres(&migration.PostgresQueryOpt { Query: "DROP DATABASE IF EXISTS " + dbName, Base: true })
	migration.RunPostgres(&migration.PostgresQueryOpt { Query: "CREATE DATABASE " + dbName, Base: true })
}
