package main

import (
	"github.com/danielbintar/angel/server-library/pubsub"
)

func main() {
	pubsub.Subscribe([]string{"angel_users_model-log"}, func(message []byte) {

	})
}
