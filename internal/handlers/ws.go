package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WSType string

const (
	TEAM     WSType = "TEAM"
	GUESS    WSType = "GUESS"
	VALIDATE WSType = "VALIDATE"
)

type WSMessage struct {
	Type    WSType          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}

	go handleConnection(conn)
}

func handleConnection(conn *websocket.Conn) {
	defer conn.Close()

	var msg WSMessage
	for {
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			return
		}
		switch {
		case msg.Type == TEAM:
			// select new team
		case msg.Type == GUESS:
			// incoming guess
		case msg.Type == VALIDATE:
			// validate
		}
		if err := conn.WriteMessage(websocket.TextMessage, msg.Payload); err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}
