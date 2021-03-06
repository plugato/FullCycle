package kafka2

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/ancalagon/simulador/application/route"
	"github.com/ancalagon/simulador/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func Producer(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route.NewRoute()

	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()

	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}

	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KAFKA_PRODUCE_TOPIC"), producer)
		time.Sleep(time.Millisecond * 500)
	}

}
