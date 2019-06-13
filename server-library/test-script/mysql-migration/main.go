package main

import (
	"os"

	"github.com/danielbintar/angel/server-library/migration"
)

func main() {
	key := "TEST_MYSQL_DATABASE"
	dbName := os.Getenv(key)

	migration.Run(&migration.QueryOpt { Query: "DROP DATABASE IF EXISTS " + dbName, Base: true })
	migration.Run(&migration.QueryOpt { Query: "CREATE DATABASE IF NOT EXISTS " + dbName, Base: true })
}
