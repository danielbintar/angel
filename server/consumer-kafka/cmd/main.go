package main

import (
	"os"

	"github.com/danielbintar/angel/server/consumer-kafka/service"
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

	_, err := service.LoadConfig(form)
	if err != nil {
		panic(err.Error)
	}
}
