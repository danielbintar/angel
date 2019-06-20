package migration

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"os"
)

// parameter to run postgres query
// 1. Query
//  The Query
//  ex: `DROP DATABASE IF EXISTS databaseName`
// 2. Base
//  Default is false
//  if false, it will run database level query
//  database will be selected from environment configuration
//    POSTGRES_DATABASE
type PostgresQueryOpt struct {
	Query string
	Base  bool
}

// Run Postgres Query based on environment configuration
// POSTGRES_USER
// POSTGRES_PASSWORD
// POSTGRES_HOST
// POSTGRES_PORT
// POSTGRES_DATABASE
func RunPostgres(opt *PostgresQueryOpt) {
	prefix := ""
	if os.Getenv("ENVIRONMENT") == "test" { prefix += "TEST_" }

	username := os.Getenv(prefix + "POSTGRES_USER")
	password := os.Getenv(prefix + "POSTGRES_PASSWORD")
	host := os.Getenv(prefix + "POSTGRES_HOST")
	port := os.Getenv(prefix + "POSTGRES_PORT")

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/", username, password, host, port)
	if opt.Base == false {
		url += os.Getenv(prefix + "POSTGRES_DATABASE")
	}
	url += "?sslmode=disable"

	db, _ := sql.Open("postgres", url)
	err := db.Ping()
	if err != nil { panic(err) }
	defer db.Close()

	_, err = db.Exec(opt.Query)
	if err != nil { panic(err) }
}
