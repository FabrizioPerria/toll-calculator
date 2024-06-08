package producers

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fabrizioperria/toll/shared/types"
)

type KafkaProducer struct {
	producer *kafka.Producer
}

var topic = "obuData"

func NewKafkaProducer() (DataProducer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
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

	return &KafkaProducer{producer: producer}, nil
}

func (kp *KafkaProducer) Produce(obuData types.OBUData) error {
	marshalData, err := json.Marshal(obuData)
	if err != nil {
		return err
	}
	return kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          marshalData,
	}, nil)
}
