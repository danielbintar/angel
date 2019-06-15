package migration

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"os"
)

// parameter to run mysql query
// 1. Query
//  The Query
//  ex: `DROP DATABASE IF EXISTS databaseName`
// 2. Base
//  Default is false
//  if false, it will run database level query
//  database will be selected from environment configuration
//    MYSQL_DATABASE
type MySQLQueryOpt struct {
	Query string
	Base  bool
}

// Run MySQL Query based on environment configuration
// MYSQL_USER
// MYSQL_PASSWORD
// MYSQL_HOST
// MYSQL_PORT
// MYSQL_DATABASE
func RunMySQL(opt *MySQLQueryOpt) {
	prefix := ""
	if os.Getenv("ENVIRONMENT") == "test" { prefix += "TEST_" }

	username := os.Getenv(prefix + "MYSQL_USER")
	password := os.Getenv(prefix + "MYSQL_PASSWORD")
	host := os.Getenv(prefix + "MYSQL_HOST")
	port := os.Getenv(prefix + "MYSQL_PORT")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, host, port)
	if opt.Base == false {
		url += os.Getenv(prefix + "MYSQL_DATABASE")
	}

	db, _ := sql.Open("mysql", url)
	err := db.Ping()
	if err != nil { panic(err) }
	defer db.Close()

	_, err = db.Exec(opt.Query)
	if err != nil { panic(err) }
}
