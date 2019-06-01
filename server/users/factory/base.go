package factory

import (
	"errors"

	"github.com/danielbintar/angel/server/users"
	"github.com/danielbintar/angel/server/users/db"
	"github.com/danielbintar/angel/server/users/model"

	"github.com/danielbintar/angel/server-library/slice"
)

func MockBase(options ...string) *users.UserManager {
	database := MockDatabase(options...)
	m := users.Instance(database)
	return m
}

func MockDatabase(options ...string) db.DatabaseManagerInterface {
	if slice.InStrings("real_database", options) {
		database := db.NewDB()
		return database
	}

	return DummyDatabase{Options: options}
}

type DummyDatabase struct {
	Options []string
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
