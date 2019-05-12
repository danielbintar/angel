package db

import (
	"os"
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
 	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbInstance *gorm.DB
	once sync.Once
)

func DB() *gorm.DB {
	once.Do(func() {
		prefix := ""
		if os.Getenv("ENVIRONMENT") == "test" { prefix += "TEST_" }
		username := os.Getenv(prefix + "MYSQL_USER")
		password := os.Getenv(prefix + "MYSQL_PASSWORD")
		host := os.Getenv(prefix + "MYSQL_HOST")
		port := os.Getenv(prefix + "MYSQL_PORT")
		dbName := os.Getenv(prefix + "MYSQL_DATABASE")
		link := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)

		dbInstance, _ = gorm.Open("mysql", link)
	})
	return dbInstance
}
