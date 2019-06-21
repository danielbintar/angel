package database

import (
	"fmt"
	"os"

	"database/sql"
	_ "github.com/lib/pq"
)

// open new Postgres connection
// will panic if not able to connect based on environment configuration
//  MYSQL_USER
//  MYSQL_PASSWORD
//  MYSQL_HOST
//  MYSQL_PORT
//  MYSQL_DATABASE
func NewPostgres() *sql.DB {
	prefix := ""
	if os.Getenv("ENVIRONMENT") == "test" {
		prefix += "TEST_"
	}
	username := os.Getenv(prefix + "POSTGRES_USER")
	password := os.Getenv(prefix + "POSTGRES_PASSWORD")
	host := os.Getenv(prefix + "POSTGRES_HOST")
	port := os.Getenv(prefix + "POSTGRES_PORT")
	dbName := os.Getenv(prefix + "POSTGRES_DATABASE")

	db, _ := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbName))

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
