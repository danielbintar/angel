package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/danielbintar/angel/server/consumer-kafka/model"
	"github.com/danielbintar/angel/server/consumer-kafka/service"

	"github.com/Shopify/sarama"
)

func main() {
	microName := os.Getenv("MICRO")
	if microName == "" {
		panic("MICRO not set")
	}

	consumerName := os.Getenv("CONSUMER")
	if consumerName == "" {
		panic("CONSUMER not set")
	}

	form := service.LoadConfigForm{
		MicroName:    microName,
		ConsumerName: consumerName,
	}

	configI, serviceErr := service.LoadConfig(form)
	if serviceErr != nil {
		panic(serviceErr.Error)
	}

	byteData, _ := json.Marshal(configI)
	var config model.Config
	json.Unmarshal(byteData, &config)

	version, err := sarama.ParseKafkaVersion("2.2.1")
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = version
	saramaConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer := Consumer{
		ready: make(chan bool, 0),
	}

	client, err := sarama.NewConsumerGroup(strings.Split(os.Getenv("KAFKA_BROKERS"), ","), "angel-consumer-kafka", saramaConfig)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}
	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			if err := client.Consume(ctx, realTopics(config.Topics), &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}

			if ctx.Err() != nil {
				return
			}

			consumer.ready = make(chan bool, 0)
		}
	}()

	<-consumer.ready
	log.Println("Kafka consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}

	cancel()
	wg.Wait()

	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

func realTopics(topics []string) []string {
	var converteds []string
	for _, topic := range topics {
		converteds = append(converteds, "angel_"+topic)
	}
	return converteds
}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready chan bool
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		session.MarkMessage(message, "")
	}

	return nil
}
