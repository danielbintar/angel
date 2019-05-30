package db

import (
	"errors"
	"fmt"
	"os"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/danielbintar/angel/server/users/model"
)

type DatabaseManagerInterface interface {
	FindUserByUsername(string) (*model.User, error)
	InsertUser(*model.User) error
}

func NewDB() DatabaseManagerInterface {
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

	db.SetConnMaxLifetime(0)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &DatabaseManager { DB: db }
}

type DatabaseManager struct {
	DB *sql.DB
}

func (self *DatabaseManager) FindUserByUsername(username string) (*model.User, error) {
	return nil, nil
}

func (self *DatabaseManager) InsertUser(user *model.User) error {
	if user == nil {
		return errors.New("cant insert nil user")
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	row, err := self.DB.Exec("INSERT INTO users(username, password, created_at, updated_at) VALUES(?, ?, ?, ?)", user.Username, user.Password, user.CreatedAt, user.UpdatedAt)

	if err != nil { return err }

	id, _ := row.LastInsertId()
	user.ID = uint(id)

	return nil
}
