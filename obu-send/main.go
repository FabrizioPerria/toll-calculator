package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/fabrizioperria/toll/shared/types"
	"github.com/gorilla/websocket"
)

func generateOBU() types.OBUData {
	latitude := rand.Float64()*180 - 90
	longitude := rand.Float64()*360 - 180
	obuID := rand.Intn(20)

	return types.OBUData{
		OBUID:     obuID,
		Latitude:  latitude,
		Longitude: longitude,
		Timestamp: time.Now().Unix(),
	}
}

func Connect(url *string) (*websocket.Conn, error) {
	var conn *websocket.Conn
	var err error
	for {
		conn, _, err = websocket.DefaultDialer.Dial(*url, nil)
		time.Sleep(1 * time.Second)
		if err == nil && conn != nil {
			break
		}
	}
	return conn, nil
}

func main() {
	url := flag.String("url", "ws://localhost:8080/obu", "url to connect to")
	conn, _ := Connect(url)
	defer conn.Close()

	for {
		for i := 0; i < rand.Intn(100); i++ {
			if err := conn.WriteJSON(generateOBU()); err != nil {
				log.Printf("Broken Pipe\n")
				// reconnect
				conn, _ = Connect(url)
			}
		}
		sleepTime := time.Duration(rand.Intn(10000))
		time.Sleep(sleepTime * time.Millisecond)
	}
}

func init() {
	rand.NewSource(time.Now().UnixNano())
}
