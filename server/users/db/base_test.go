package db_test

import (
	"testing"
	"time"

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
	assert.NotNil(t, database.InsertUser(nil))

	u := &model.User{}
	assert.Equal(t, uint(0), u.ID)
	assert.Equal(t, time.Time{}, u.CreatedAt)
	assert.Equal(t, time.Time{}, u.UpdatedAt)

	assert.Nil(t, database.InsertUser(u))

	assert.NotEqual(t, uint(0), u.ID)
	assert.NotEqual(t, time.Time{}, u.CreatedAt)
	assert.NotEqual(t, time.Time{}, u.UpdatedAt)
	assert.Equal(t, u.CreatedAt, u.UpdatedAt)
}

func TestFindUserByUsername(t *testing.T) {
	gotenv.Load("../.env")
	database := db.NewDB()
	u, err := database.FindUserByUsername("asd")
	assert.Nil(t, u)
	assert.Nil(t, err)

	database.InsertUser(&model.User{Username: "lala"})

	u, err = database.FindUserByUsername("lala")
	assert.NotNil(t, u)
	assert.Nil(t, err)
}
