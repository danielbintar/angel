package main

import (
	"github.com/danielbintar/angel/server-library/pubsub"
)

func main() {
	pubsub.Subscribe("angel_users_model-log", handle)
}

func handle(message []byte) {

}
