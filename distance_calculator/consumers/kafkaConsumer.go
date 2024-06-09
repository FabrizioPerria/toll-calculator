package consumers

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fabrizioperria/toll/distance_calculator/service"
	constants "github.com/fabrizioperria/toll/shared"
	"github.com/fabrizioperria/toll/shared/types"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
	service.Calculator
}

func NewKafkaConsumer() (DataConsumer, error) {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          1,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}
	kafkaConsumer.SubscribeTopics([]string{constants.KafkaObuDataTopic}, nil)

	var srvc service.Calculator
	srvc = service.NewCalculatorService()
	srvc = service.NewLogServiceMiddleware(srvc)
	return &KafkaConsumer{
		consumer:   kafkaConsumer,
		Calculator: srvc,
	}, nil
}

func (kc *KafkaConsumer) Consume() (types.OBUData, error) {
	message, err := kc.consumer.ReadMessage(-1)
	if err != nil {
		return types.OBUData{}, err
	}

	obuData := types.OBUData{}
	if err = json.Unmarshal(message.Value, &obuData); err != nil {
		return types.OBUData{}, err
	}
	if _, err := kc.Distance(obuData); err != nil {
		return obuData, err
	}
	return obuData, nil
}
