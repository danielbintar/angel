package db

import (
	"errors"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/danielbintar/angel/server-library/database"
	"github.com/danielbintar/angel/server/users/model"
)

type DatabaseManagerInterface interface {
	Close() error
	FindUserByUsername(string) (*model.User, error)
	InsertUser(*model.User) error
}

func NewDB() DatabaseManagerInterface {
	return &DatabaseManager{DB: database.NewMySQL()}
}

type DatabaseManager struct {
	DB *sql.DB
}

func (self *DatabaseManager) Close() error {
	return self.DB.Close()
}

func (self *DatabaseManager) FindUserByUsername(username string) (*model.User, error) {
	record := self.DB.QueryRow("SELECT id, username, password, created_at, updated_at FROM users WHERE username=? LIMIT 1", username)

	var user model.User
	err := record.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (self *DatabaseManager) InsertUser(user *model.User) error {
	if user == nil {
		return errors.New("cant insert nil user")
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	row, err := self.DB.Exec("INSERT INTO users(username, password, created_at, updated_at) VALUES(?, ?, ?, ?)", user.Username, user.Password, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	id, _ := row.LastInsertId()
	user.ID = uint(id)

	return nil
}
