package consumers

import (
	"encoding/json"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fabrizioperria/toll/aggregator/client"
	"github.com/fabrizioperria/toll/distance_calculator/service"
	constants "github.com/fabrizioperria/toll/shared"
	"github.com/fabrizioperria/toll/shared/types"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
	service.Calculator
	client *client.HTTPAggregatorClient
}

func NewKafkaConsumer(server string) (DataConsumer, error) {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          1,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}
	kafkaConsumer.SubscribeTopics([]string{constants.KafkaObuDataTopic}, nil)

	return NewLogConsumerMiddleware(&KafkaConsumer{
		consumer:   kafkaConsumer,
		Calculator: service.NewCalculatorService(),
		client:     client.NewHTTPAggregatorClient(constants.AggregatorHttpClient),
	}), nil
}

func (kc *KafkaConsumer) Consume() (types.OBUData, error) {
	message, err := kc.consumer.ReadMessage(-1)
	if err != nil {
		return types.OBUData{}, nil
	}

	obuData := types.OBUData{}
	if err = json.Unmarshal(message.Value, &obuData); err != nil {
		return types.OBUData{}, err
	}
	rawDistance, err := kc.Distance(obuData)
	if err != nil {
		return obuData, err
	}
	distance := types.Distance{
		OBUID:     obuData.OBUID,
		Value:     rawDistance,
		Timestamp: time.Now().Unix(),
	}
	kc.client.Aggregate(distance)
	return obuData, nil
}
