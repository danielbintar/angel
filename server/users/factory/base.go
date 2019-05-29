package factory

import (
	"fmt"
	"errors"

	"github.com/danielbintar/angel/server/users"
	"github.com/danielbintar/angel/server/users/db"
	"github.com/danielbintar/angel/server/users/model"

	"github.com/danielbintar/angel/server-library/slice"
)

func MockBase() *users.UserManager {
	database := MockDatabase()
	m := users.Instance(database)
	return m
}

func MockDatabase(options ...string) db.DatabaseManagerInterface {
	if slice.InStrings("broken_database", options) {
		fmt.Println("hmmm")
	}

	return DummyDatabase{}
}

type DummyDatabase struct {}
func (self DummyDatabase) InsertUser(user model.User) error { return nil }

type BrokenDatabase struct {}
func (self BrokenDatabase) InsertUser(user model.User) error { return errors.New("broken") }
