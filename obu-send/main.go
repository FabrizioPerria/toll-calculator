package main

import (
	"flag"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/fabrizioperria/toll/shared/types"
	"github.com/gorilla/websocket"
)

func generateOBU() types.OBUData {
	latitude := rand.Float64()*180 - 90
	longitude := rand.Float64()*360 - 180
	obuID := strconv.Itoa(rand.Intn(20))

	obuData := types.OBUData{
		ObuId:     obuID,
		Latitude:  latitude,
		Longitude: longitude,
		Timestamp: time.Now().Unix(),
	}

	log.Printf("Generated OBU data: %+v\n", obuData)
	return obuData
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
	url := flag.String("url", "ws://localhost:8084/obu", "url to connect to")
	flag.Parse()
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
		sleepTime := time.Duration(rand.Intn(1000))
		time.Sleep(sleepTime * time.Millisecond)
	}
}

func init() {
	rand.NewSource(time.Now().UnixNano())
}
