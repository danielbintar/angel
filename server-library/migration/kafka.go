package migration

import (
	"os"
	"strings"

	"github.com/Shopify/sarama"
)

// create kafka topic, will panic if already created
// ex: migration.CreateKafkaTopic("increase-sold-count", 3, 1)
// topic name should be hyphen case (-)
func CreateKafkaTopic(topic string, numPartition int32, replicationFactor int16) {
	handleClusterAdminKafka(func(admin sarama.ClusterAdmin) {
		err := admin.CreateTopic(realTopic(topic), &sarama.TopicDetail{
			NumPartitions:     numPartition,
			ReplicationFactor: replicationFactor,
		}, false)
		if err != nil {
			panic(err)
		}
	})
}

// delete kafka topic, will do nothing if not exists
// ex: migration.DeleteKafkaTopic("increase-sold-count")
func DeleteKafkaTopic(topic string) {
	handleClusterAdminKafka(func(admin sarama.ClusterAdmin) {
		err := admin.DeleteTopic(realTopic(topic))
		if err != nil && err.Error() != "kafka server: Request was for a topic or partition that does not exist on this broker." {
			panic(err)
		}
	})
}

func realTopic(topic string) string {
	prefix := ""
	if os.Getenv("ENVIRONMENT") == "test" {
		prefix = "TEST_"
	}
	return prefix + "angel_" + topic
}

func handleClusterAdminKafka(handle func(admin sarama.ClusterAdmin)) {
	prefix := ""
	if os.Getenv("ENVIRONMENT") == "test" {
		prefix = "TEST_"
	}

	brokerAddrs := strings.Split(os.Getenv(prefix+"KAFKA_BROKERS"), ",")

	config := sarama.NewConfig()
	version, err := sarama.ParseKafkaVersion("2.2.1")
	if err != nil {
		panic(err)
	}
	config.Version = version

	admin, err := sarama.NewClusterAdmin(brokerAddrs, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := admin.Close()
		if err != nil {
			panic(err)
		}
	}()

	handle(admin)
}
