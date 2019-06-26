package handler

import (
	"errors"

	"github.com/Shopify/sarama"
)

func Handle(message sarama.ConsumerMessage) error {
	if message.Timestamp.IsZero() {
		return errors.New("invalid timestamp")
	}

	return nil
}
