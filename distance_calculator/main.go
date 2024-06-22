package main

import (
	"log"
	"os"

	"github.com/fabrizioperria/toll/distance_calculator/consumers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	server := os.Getenv("KAFKA_SERVER")
	var kafkaConsumer consumers.DataConsumer
	kafkaConsumer, err := consumers.NewKafkaConsumer(server)
	if err != nil {
		panic(err)
	}

	for {
		_, err := kafkaConsumer.Consume()
		if err != nil {
			log.Printf("Error consuming message\n")
		}
	}
}
