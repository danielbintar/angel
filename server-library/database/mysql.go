package database

import (
	"fmt"
	"os"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewMySQL() *sql.DB {
	prefix := ""
	if os.Getenv("ENVIRONMENT") == "test" { prefix += "TEST_" }
	username := os.Getenv(prefix + "MYSQL_USER")
	password := os.Getenv(prefix + "MYSQL_PASSWORD")
	host := os.Getenv(prefix + "MYSQL_HOST")
	port := os.Getenv(prefix + "MYSQL_PORT")
	dbName := os.Getenv(prefix + "MYSQL_DATABASE")

	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", username, password, host, port, dbName))

	err := db.Ping()
	if err != nil { panic(err) }

	return db
}