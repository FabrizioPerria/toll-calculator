package producers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fabrizioperria/toll/shared/types"
)

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer() (DataProducer, error) {
	kafkaServer := os.Getenv("KAFKA_SERVER")
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServer,
	})
	if err != nil {
		return nil, err
	}

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return NewLogMiddleware(&KafkaProducer{producer: producer}), nil
}

func (kp *KafkaProducer) Produce(obuData types.OBUData) error {
	marshalData, err := json.Marshal(obuData)
	if err != nil {
		return err
	}
	topic := os.Getenv("KAFKA_TOPIC")
	fmt.Println("Producing to topic: ", topic)
	return kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          marshalData,
	}, nil)
}

func (kp *KafkaProducer) Flush() int {
	return kp.producer.Flush(15 * 1000)
}

func (kp *KafkaProducer) Close() {
	kp.producer.Close()
}
