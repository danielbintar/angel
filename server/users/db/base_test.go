package db_test

import (
	"testing"

	"github.com/danielbintar/angel/server/users/db"

	"github.com/subosito/gotenv"

	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	assert.NotNil(t, db.DB())
	gotenv.Load("../.env")
	assert.NotNil(t, db.DB())
}
