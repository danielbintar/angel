package database_test

import (
	"os"
	"testing"

	"github.com/danielbintar/angel/server-library/database"

	"github.com/stretchr/testify/assert"
)

func TestNewMySQL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		assert.NotPanics(t, func() { database.NewMySQL() })	
	})

	t.Run("wrong configuration", func(t *testing.T) {
		username := os.Getenv("TEST_MYSQL_USER")
		os.Setenv("TEST_MYSQL_USER", "invaliduser1231")
		assert.Panics(t, func() { database.NewMySQL() })	
		os.Setenv("TEST_MYSQL_USER", username)
	})
}
