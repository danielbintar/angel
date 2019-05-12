package main

import (
	"os"

	"github.com/danielbintar/angel/server/users/db/migration"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()

	var query string
	if len(os.Args) == 2 && os.Args[1] == "down" {
		query = "DROP TABLE IF EXISTS users"
	} else {
		query = `CREATE TABLE IF NOT EXISTS users (
		id INT NOT NULL AUTO_INCREMENT,
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL DEFAULT NULL,
		PRIMARY KEY (id)
		)  ENGINE=INNODB;`
	}

	migration.Run(&migration.QueryOpt { Query: query })
}
