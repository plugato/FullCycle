package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type kafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

func NewKafkaConsumer(msgChan chan *ckafka.Message) *kafkaConsumer {
	return &kafkaConsumer{
		MsgChan: msgChan,
	}
}

func (k *kafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKER"),
		"group.id":          os.Getenv("KAFKA_GROUP_ID"),
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}

	topics := []string{os.Getenv("KAFKA_READ_TOPIC")}
	c.SubscribeTopics(topics, nil)
	fmt.Println("Subscribed to topic", topics)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MsgChan <- msg
		}
	}
}
