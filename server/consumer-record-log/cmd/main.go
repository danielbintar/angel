package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/danielbintar/angel/server-library/model"

	"github.com/Shopify/sarama"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()

	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	if err != nil { panic(err) }
	defer func() { if err := consumer.Close(); err != nil { panic(err) } }()

	partConsumer, err := consumer.ConsumePartition("angel_users_model-log", 0, sarama.OffsetNewest)
	if err != nil { panic(err) }
	defer func() { if err := partConsumer.Close(); err != nil { panic(err) } }()

	fmt.Println("listen to kafka topic angel_users.model-log")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0

	Loop:
	for {
		select {
		case msg := <- partConsumer.Messages():
			fmt.Printf("Consumed message: [%s], offset: [%d]\n", msg.Value, msg.Offset)
			var payload model.RequestPayload
			json.Unmarshal(msg.Value, &payload)
			fmt.Println(payload)
			consumed++
		case <-signals:
			break Loop
		}
	}

	fmt.Printf("Consumed: %d\n", consumed)
}
