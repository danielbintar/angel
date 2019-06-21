package db

import (
	"errors"

	"database/sql"

	"github.com/danielbintar/angel/server-library/database"
	"github.com/danielbintar/angel/server/users/model"
)

type DatabaseManagerInterface interface {
	Close() error
	FindUserByUsername(string) (*model.User, error)
	InsertUser(*model.User) error
}

func NewDB() DatabaseManagerInterface {
	return &DatabaseManager{DB: database.NewPostgres()}
}

type DatabaseManager struct {
	DB *sql.DB
}

func (self *DatabaseManager) Close() error {
	return self.DB.Close()
}

func (self *DatabaseManager) FindUserByUsername(username string) (*model.User, error) {
	record := self.DB.QueryRow(`SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1 LIMIT 1`, username)

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

	return self.DB.QueryRow(
		`INSERT INTO users(username, password) VALUES($1, $2) RETURNING id, created_at, updated_at`,
		user.Username, user.Password).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}
