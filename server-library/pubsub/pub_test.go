package pubsub_test

import (
	"os"
	"testing"

	"github.com/danielbintar/angel/server-library/pubsub"

	"github.com/stretchr/testify/assert"
)

func TestNewKafkaAsyncProducer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		assert.NotPanics(t, func() { pubsub.NewKafkaAsyncProducer() })
		assert.NotNil(t, pubsub.NewKafkaAsyncProducer())
	})

	t.Run("wrong configuration", func(t *testing.T) {
		brokers := os.Getenv("TEST_KAFKA_BROKERS")
		os.Setenv("TEST_KAFKA_BROKERS", "invalid address")
		assert.Panics(t, func() { pubsub.NewKafkaAsyncProducer() })
		os.Setenv("TEST_KAFKA_BROKERS", brokers)
	})
}

func TestClose(t *testing.T) {
	publisher := pubsub.NewKafkaAsyncProducer()
	assert.NotPanics(t, func() { publisher.Close() })
}

func TestPublish(t *testing.T) {
	publisher := pubsub.NewKafkaAsyncProducer()
	assert.NotPanics(t, func() { publisher.Publish("zzzza", "message") })
}
