package database_test

import (
	"os"
	"testing"

	"github.com/danielbintar/angel/server-library/database"

	"github.com/stretchr/testify/assert"
)

func TestNewPostgres(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		assert.NotPanics(t, func() { database.NewPostgres() })
	})

	t.Run("wrong configuration", func(t *testing.T) {
		username := os.Getenv("TEST_POSTGRES_USER")
		os.Setenv("TEST_POSTGRES_USER", "invaliduser1231")
		assert.Panics(t, func() { database.NewPostgres() })
		os.Setenv("TEST_POSTGRES_USER", username)
	})
}
