package main

import (
	"log"

	"github.com/fabrizioperria/toll/distance_calculator/consumers"
)

func main() {
	kafkaConsumer, err := consumers.NewKafkaConsumer()
	if err != nil {
		panic(err)
	}

	kafkaConsumer = consumers.NewLogMiddleware(kafkaConsumer)

	for {
		_, err := kafkaConsumer.Consume()
		if err != nil {
			log.Printf("Error consuming message %v\n", err)
		}
	}
}
