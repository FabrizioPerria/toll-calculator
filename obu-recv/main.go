package main

import (
	"fmt"
	"net/http"

	"github.com/fabrizioperria/toll/obu-recv/producers"
	"github.com/fabrizioperria/toll/shared/types"
	"github.com/gorilla/websocket"
)

type dataReceiver struct {
	msg      chan types.OBUData
	producer producers.DataProducer
	conn     *websocket.Conn
}

func newDataReceiver() *dataReceiver {
	producer, err := producers.NewKafkaProducer()
	if err != nil {
		panic(err)
	}
	producer = producers.NewLogMiddleware(producer)
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
		if err := dr.producer.Produce(obuData); err != nil {
			fmt.Printf("%+v\n", err)
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

func main() {
	dataReceiver := newDataReceiver()
	// defer dataReceiver.producer.Flush(15 * 1000)
	// defer dataReceiver.producer.Close()

	http.HandleFunc("/obu", dataReceiver.obuHandler)
	http.ListenAndServe(":8080", nil)
}
