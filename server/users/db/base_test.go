package db_test

import (
	"testing"
	"time"

	"github.com/danielbintar/angel/server/users/db"
	"github.com/danielbintar/angel/server/users/factory"
	"github.com/danielbintar/angel/server/users/model"

	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {
	assert.NotPanics(t, func() { db.NewDB() })
}

func TestClose(t *testing.T) {
	database := db.NewDB()
	_, err := database.FindUserByUsername("lala")
	assert.Nil(t, err)

	database.Close()
	_, err = database.FindUserByUsername("lala")
	assert.NotNil(t, err)
}

func TestInsertUser(t *testing.T) {
	database := db.NewDB()

	t.Run("nil user", func(t *testing.T) {
		assert.NotNil(t, database.InsertUser(nil))
	})

	t.Run("non nil user", func(t *testing.T) {
		u := &model.User{}
		assert.Equal(t, uint(0), u.ID)
		assert.Equal(t, time.Time{}, u.CreatedAt)
		assert.Equal(t, time.Time{}, u.UpdatedAt)

		assert.Nil(t, database.InsertUser(u))

		assert.NotEqual(t, uint(0), u.ID)
		assert.NotEqual(t, time.Time{}, u.CreatedAt)
		assert.NotEqual(t, time.Time{}, u.UpdatedAt)
		assert.Equal(t, u.CreatedAt, u.UpdatedAt)
	})

	t.Run("broken database", func(t *testing.T) {
		database := factory.MockDatabase("broken_real_database")
		 u := model.User{}
		assert.NotNil(t, database.InsertUser(&u))
	})
}

func TestFindUserByUsername(t *testing.T) {
	database := db.NewDB()

	t.Run("not exists user", func(t *testing.T) {
		u, err := database.FindUserByUsername("asd")
		assert.Nil(t, u)
		assert.Nil(t, err)
	})

	t.Run("exists user", func(t *testing.T) {
		database.InsertUser(&model.User{Username: "lala"})

		u, err := database.FindUserByUsername("lala")
		assert.NotNil(t, u)
		assert.Nil(t, err)
	})

	t.Run("broken database", func(t *testing.T) {
		mockDatabase := factory.MockDatabase("broken_real_database")
		database.InsertUser(&model.User{Username: "lala"})
		u, err := mockDatabase.FindUserByUsername("asd")
		assert.Nil(t, u)
		assert.NotNil(t, err)
	})
}
