package main

import (
	"io/ioutil"
	"os"
)

func main() {
	microName := os.Getenv("MICRO")
	if microName == "" { panic("MICRO not set") }

	consumerName := os.Getenv("CONSUMER")
	if consumerName == "" { panic("CONSUMER not set") }

	_, err := ioutil.ReadFile("consumers/" + microName + "/" + consumerName + "/config.yaml")
	if err != nil { panic(consumerName + " is not found in " + microName) }
}
