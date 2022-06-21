package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProducer() *ckafka.Producer {

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKER"),
	}
	fmt.Println("KAFKA_BROKER", os.Getenv("KAFKA_BROKER"))
	p, err := ckafka.NewProducer(configMap)
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
	return p

}

func Publish(msg string, topic string, producer *ckafka.Producer) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}
	err := producer.Produce(message, nil)
	if err != nil {
		return err
	}
	return nil
}
