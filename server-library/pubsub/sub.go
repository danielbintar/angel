package pubsub

import (
	"fmt"
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Shopify/sarama"
)

func Subscribe(topics []string, handle func(message []byte)) {
	if len(topics) == 0 { panic("no topic to subscribed") }

	config := sarama.NewConfig()
	config.Version = sarama.V2_2_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer := Consumer{ handle: handle }

	ctx := context.Background()
	prefix := ""
	if os.Getenv("ENVIRONMENT") == "test" { prefix = "TEST_" }
	brokerAddrs := strings.Split(os.Getenv(prefix + "KAFKA_BROKERS"), ",")
	client, err := sarama.NewConsumerGroup(brokerAddrs, topics[0], config)
	if err != nil { panic(err) }

	consumer.ready = make(chan bool, 0)

	go func() {
		for {
			err := client.Consume(ctx, realTopics(topics), &consumer)
			if err != nil { panic(err) }
		}
	}()

	<-consumer.ready

	fmt.Println("Consuming " + strings.Join(topics, ", "))

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	<-sigterm
}

func realTopics(topics []string) []string {
	var r []string
	for _, topic := range topics {
		r = append(r, realTopic(topic))
	}
	return r
}

func realTopic(topic string) string {
	prefix := ""
	if os.Getenv("ENVIRONMENT") == "test" { prefix = "TEST_" }
	return prefix + "angel_" + topic
}

type Consumer struct {
	ready chan bool
	handle func(message []byte)
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (self *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(self.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (self *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (self *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		self.handle(message.Value)
		fmt.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		session.MarkMessage(message, "")
	}

	return nil
}