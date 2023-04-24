package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const wsServerEndpoint = "ws://localhost:40000/websocket"

func GameServer() {
	dialer := websocket.Dialer{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, resp, err := dialer.Dial(wsServerEndpoint, nil)

	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}

	log.Println("Connected to WebSocket server. Response status:", resp.Status)

	defer conn.Close()

	for {
		uuid := uuid.New()

		err := conn.WriteMessage(websocket.TextMessage, []byte("Hello, world! "+uuid.String()))
		if err != nil {
			log.Println("Error writing to WebSocket:", err)
			break
		}

		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading from WebSocket:", err)
			break
		}

		log.Printf("Received message: %s", message)

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error writing to WebSocket:", err)
			break
		}
	}

}

func main() {
	GameServer()
}
