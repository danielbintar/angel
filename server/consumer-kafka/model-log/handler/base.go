package handler

import (
	"github.com/danielbintar/angel/server/consumer-kafka/model-log/model"

	"github.com/Shopify/sarama"
)

func Handle(message sarama.ConsumerMessage) error {
	if message.Timestamp.IsZero() {
		return model.NewErrInvalidMessage("invalid timestamp")
	}

	return nil
}
