package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/JaanLavaerts/ticktickbrick/internal/models"
	"github.com/JaanLavaerts/ticktickbrick/internal/util"
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

	client := &models.Client{
		User: models.User{
			Id:       util.GenerateID(),
			Username: "guest",
			Lives:    3,
		},
		Conn: conn,
		Send: make(chan []byte),
	}

	go handleConnection(client)
}

//TODO: split into write and read handlers

func handleConnection(client *models.Client) {
	defer client.Conn.Close()

	var msg WSMessage
	for {
		err := client.Conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			return
		}
		switch {
		case msg.Type == TEAM:
			fmt.Println("TEAM")
		case msg.Type == GUESS:
			// incoming guess
		case msg.Type == VALIDATE:
			// validate
		}
		if err := client.Conn.WriteMessage(websocket.TextMessage, msg.Payload); err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}
