package model_test

import (
	"testing"

	"github.com/danielbintar/angel/server/users/model"

	"github.com/stretchr/testify/assert"
)

func TestSerialize(t *testing.T) {
	var nilU *model.User
	u := &model.User{}

	assert.NotNil(t, u.Serialize())
	assert.Nil(t, nilU.Serialize())
}
