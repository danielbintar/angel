package model_test

import (
	"testing"

	"github.com/danielbintar/angel/server/consumer-kafka/model-log/model"

	"github.com/stretchr/testify/assert"
)

func TestNewErrInvalidMessage(t *testing.T) {
	message := "invalid timestamp"
	err := model.NewErrInvalidMessage(message)
	assert.NotNil(t, err)
}

func TestErrInvalidMessageError(t *testing.T) {
	message := "invalid timestamp"
	err := model.NewErrInvalidMessage(message)
	assert.Equal(t, message, err.Error())
}
