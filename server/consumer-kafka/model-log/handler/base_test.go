package handler_test

import (
	"testing"

	"github.com/danielbintar/angel/server/consumer-kafka/model-log/handler"	

	"github.com/Shopify/sarama"

	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {
	assert.NotPanics(t, func() { handler.Handle(sarama.ConsumerMessage{}) })
}
