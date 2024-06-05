package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fabrizioperria/toll/shared/types"
	"github.com/gorilla/websocket"
)

type dataReceiver struct {
	msg      chan types.OBUData
	producer *kafka.Producer
	conn     *websocket.Conn
}

func newDataReceiver() *dataReceiver {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return &dataReceiver{
		msg:      make(chan types.OBUData, 100),
		producer: producer,
		conn:     nil,
	}
}

func (dr *dataReceiver) obuHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	if dr.conn != nil {
		dr.conn.Close()
	}
	dr.conn = conn

	go dr.recvLoop()
	go dr.produceData()
}

func (dr *dataReceiver) produceData() {
	for obuData := range dr.msg {
		marshalData, err := json.Marshal(obuData)
		if err != nil {
			return
		}
		err = dr.producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          marshalData,
		}, nil)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func (dr *dataReceiver) recvLoop() {
	defer dr.conn.Close()
	for {
		var obuData types.OBUData
		if err := dr.conn.ReadJSON(&obuData); err != nil {
			fmt.Printf("%+v\n", err)
			if websocket.IsCloseError(err, websocket.CloseAbnormalClosure) {
				break
			}
			continue
		}
		dr.msg <- obuData
	}
}

var topic = "obuData"

func main() {
	dataReceiver := newDataReceiver()
	defer dataReceiver.producer.Flush(15 * 1000)
	defer dataReceiver.producer.Close()

	http.HandleFunc("/obu", dataReceiver.obuHandler)
	http.ListenAndServe(":8080", nil)
}
