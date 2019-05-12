package migration

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"os"
)

type QueryOpt struct {
	Query string
	Base  bool
}


// TODO: BECOME LIBRARY
func Run(opt *QueryOpt) {
	prefix := ""
	if os.Getenv("ENVIRONMENT") == "TEST" { prefix += "TEST_" }

	username := os.Getenv(prefix + "MYSQL_USER")
	password := os.Getenv(prefix + "MYSQL_PASSWORD")
	host := os.Getenv(prefix + "MYSQL_HOST")
	port := os.Getenv(prefix + "MYSQL_PORT")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, host, port)
	if opt.Base == false {
		url += os.Getenv(prefix + "MYSQL_DATABASE")
	}

	db, err := sql.Open("mysql", url)
	if err != nil { panic(err) }
	defer db.Close()

	_, err = db.Exec(opt.Query)
	if err != nil { panic(err) }
}
