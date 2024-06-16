package main

import (
	"flag"
	"log"

	"github.com/fabrizioperria/toll/distance_calculator/consumers"
)

func main() {
	server := flag.String("server", "localhost", "Kafka server")
	flag.Parse()
	var kafkaConsumer consumers.DataConsumer
	kafkaConsumer, err := consumers.NewKafkaConsumer(*server)
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
