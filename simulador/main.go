package main

import (
	"fmt"
	"log"

	kafka2 "github.com/ancalagon/simulador/application/kafka"
	"github.com/ancalagon/simulador/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {

	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Producer(msg)
	}
	// producer := kafka.NewKafkaProducer()

	// kafka.Publish("ola12122", "teste", producer)

	// for {
	// 	_ = 1
	// }
	// route := route.Route{
	// 	ID:       "1",
	// 	ClientId: "1",
	// }

	// route.LoadPositions()

	// stringjson, _ := route.ExportJsonPositions()
	// fmt.Println(stringjson[1])
}
