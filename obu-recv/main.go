package main

import (
	"fmt"
	"net/http"

	"github.com/fabrizioperria/toll/shared/types"
	"github.com/gorilla/websocket"
)

type dataReceiver struct {
	msg  chan types.OBUData
	conn *websocket.Conn
}

func newDataReceiver() *dataReceiver {
	return &dataReceiver{
		msg:  make(chan types.OBUData, 100),
		conn: nil,
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
	dr.conn = conn

	go dr.recvLoop()

	go func() {
		for {
			select {
			case obuData := <-dr.msg:
				fmt.Println("Received data from OBU:", obuData)
			}
		}
	}()
}

func (dr *dataReceiver) recvLoop() {
	fmt.Println("A new client connected!")
	for {
		var obuData types.OBUData
		if err := dr.conn.ReadJSON(&obuData); err != nil {
			fmt.Println(err)
			continue
		}
		dr.msg <- obuData
	}
}

func main() {
	dataReceiver := newDataReceiver()
	defer dataReceiver.conn.Close()
	http.HandleFunc("/obu", dataReceiver.obuHandler)
	http.ListenAndServe(":8080", nil)
}
