package pubsub

import (
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

type AsyncPublisher interface {
	Publish(id string, message string)
	Close()
}

type KafkaAsyncProducer struct {
	producer sarama.AsyncProducer
}

func NewKafkaAsyncProducer() AsyncPublisher {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = false
	config.Producer.Return.Errors = false

	prefix := ""
	if os.Getenv("ENVIRONMENT") == "test" { prefix = "TEST_" }
	brokers := strings.Split(os.Getenv(prefix + "KAFKA_BROKERS"), ",")
	kafkaProducer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil { panic(err) }
	return &KafkaAsyncProducer{producer: kafkaProducer}
}

func (self *KafkaAsyncProducer) Close() {
	self.producer.AsyncClose()
}

func (self *KafkaAsyncProducer) Publish(id string, message string) {
	id = "angel_" + id

	producerMessage := &sarama.ProducerMessage{
		Topic:     id,
		Value:     sarama.StringEncoder(message),
		Key:       sarama.StringEncoder(""),
		Timestamp: time.Now(),
	}
	self.producer.Input() <- producerMessage
}
