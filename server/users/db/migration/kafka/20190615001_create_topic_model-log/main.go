package main

import (
	"os"

	"github.com/danielbintar/angel/server-library/migration"
)

func main() {
	topic := "users_model-log"
	if len(os.Args) == 2 && os.Args[1] == "down" {
		migration.DeleteKafkaTopic(topic)
	} else {
		migration.CreateKafkaTopic(topic, 3, 1)
	}
}
