package main

import (
	"os"
)

func main() {
	microName := os.Getenv("MICRO")
	if microName == "" { panic("MICRO not set") }

	consumerName := os.Getenv("CONSUMER")
	if consumerName == "" { panic("CONSUMER not set") }
}
