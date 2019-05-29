package db_test

import (
	"testing"

	"github.com/danielbintar/angel/server/users/db"
	"github.com/danielbintar/angel/server/users/model"

	"github.com/subosito/gotenv"

	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {
	assert.Panics(t, func() { db.NewDB() })
	gotenv.Load("../.env")
	assert.NotPanics(t, func() { db.NewDB() })
}

func TestInsertUser(t *testing.T) {
	gotenv.Load("../.env")
	database := db.NewDB()
	assert.Nil(t, database.InsertUser(model.User{}))
}
