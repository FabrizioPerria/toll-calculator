package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fabrizioperria/toll/obu-recv/producers"
	"github.com/fabrizioperria/toll/shared/types"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

type dataReceiver struct {
	msg      chan types.OBUData
	producer producers.DataProducer
}

func newDataReceiver() *dataReceiver {
	producer, err := producers.NewKafkaProducer()
	if err != nil {
		panic(err)
	}
	return &dataReceiver{
		msg:      make(chan types.OBUData, 100),
		producer: producer,
	}
}

func (dr *dataReceiver) obuHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error while upgrading connection: %v", err)
		return
	}

	go dr.recvLoop(conn)
	go dr.produceData()
}

func (dr *dataReceiver) produceData() {
	for obuData := range dr.msg {
		if err := dr.producer.Produce(obuData); err != nil {
			fmt.Printf("%+v\n", err)
		}
	}
}

func (dr *dataReceiver) recvLoop(conn *websocket.Conn) {
	defer conn.Close()
	for {
		var obuData types.OBUData
		if err := conn.ReadJSON(&obuData); err != nil {
			fmt.Printf("%+v\n", err)
			break
		}
		dr.msg <- obuData
	}
}

func main() {
	godotenv.Load()
	listenAddr := os.Getenv("RECEIVER_LISTEN_ADDR")
	dataReceiver := newDataReceiver()
	defer dataReceiver.producer.Flush()
	defer dataReceiver.producer.Close()

	http.HandleFunc("/obu", dataReceiver.obuHandler)
	http.ListenAndServe(listenAddr, nil)
}
