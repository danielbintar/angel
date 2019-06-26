package handler_test

import (
	"testing"
	"time"

	"github.com/danielbintar/angel/server/consumer-kafka/model-log/handler"
	"github.com/danielbintar/angel/server/consumer-kafka/model-log/model"

	"github.com/Shopify/sarama"

	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {
	t.Run("no timestamp", func(t *testing.T) {
		err := handler.Handle(sarama.ConsumerMessage{})
		assert.NotNil(t, err)

		_, ok := err.(*model.ErrInvalidMessage)
		assert.True(t, ok)
		assert.NotEqual(t, "", err.Error())
	})

	t.Run("valid message", func(t *testing.T) {
		err := handler.Handle(sarama.ConsumerMessage{
			Timestamp: time.Now(),
		})
		assert.Nil(t, err)
	})
}
