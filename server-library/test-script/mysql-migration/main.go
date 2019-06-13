package main

import (
	"os"

	"github.com/danielbintar/angel/server-library/migration"
)

func main() {
	key := "TEST_MYSQL_DATABASE"
	dbName := os.Getenv(key)

	migration.RunMySQL(&migration.MySQLQueryOpt { Query: "DROP DATABASE IF EXISTS " + dbName, Base: true })
	migration.RunMySQL(&migration.MySQLQueryOpt { Query: "CREATE DATABASE IF NOT EXISTS " + dbName, Base: true })
}
