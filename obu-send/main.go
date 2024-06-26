package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/fabrizioperria/toll/shared/types"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
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
	godotenv.Load()
	url := os.Getenv("WEBSOCKET_URL")
	conn, _ := Connect(&url)
	defer conn.Close()

	for {
		for i := 0; i < rand.Intn(1000); i++ {
			if err := conn.WriteJSON(generateOBU()); err != nil {
				log.Printf("Broken Pipe\n")
			}
		}
		sleepTime := time.Duration(rand.Intn(10))
		time.Sleep(sleepTime * time.Millisecond)
	}
}

func init() {
	rand.NewSource(time.Now().UnixNano())
}
