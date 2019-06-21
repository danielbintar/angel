package factory

import (
	"errors"
	"fmt"

	"database/sql"

	"github.com/danielbintar/angel/server/users"
	"github.com/danielbintar/angel/server/users/db"
	"github.com/danielbintar/angel/server/users/model"

	libFactory "github.com/danielbintar/angel/server-library/factory"
	"github.com/danielbintar/angel/server-library/slice"
)

func MockBase(options ...string) *users.UserManager {
	database := MockDatabase(options...)
	publisher := libFactory.MockAsyncPublisher()
	m := users.UserManager{
		DatabaseManager: database,
		Publisher:       publisher,
	}
	return &m
}

func MockDatabase(options ...string) db.DatabaseManagerInterface {
	if slice.InStrings("real_database", options) {
		database := db.NewDB()
		return database
	}

	if slice.InStrings("broken_real_database", options) {
		database := NewBrokenDB()
		return database
	}

	return DummyDatabase{Options: options}
}

func NewBrokenDB() db.DatabaseManagerInterface {
	database, _ := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "", "", "", "", ""))
	return &db.DatabaseManager{DB: database}
}

type DummyDatabase struct {
	Options []string
}

func (self DummyDatabase) Close() error {
	return nil
}

func (self DummyDatabase) InsertUser(_ *model.User) error {
	if slice.InStrings("broken_insert_user", self.Options) {
		return errors.New("broken")
	}

	return nil
}

func (self DummyDatabase) FindUserByUsername(_ string) (*model.User, error) {
	if slice.InStrings("broken_find_user_by_username", self.Options) {
		return nil, errors.New("broken")
	}

	if slice.InStrings("find_user_by_username_404", self.Options) {
		return nil, nil
	}

	return &model.User{}, nil
}
