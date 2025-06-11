package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/JaanLavaerts/ticktickbrick/internal/models"
	"github.com/gorilla/websocket"
)

type WSType string

const (
	CREATE_ROOM WSType = "CREATE_ROOM"
	JOIN_ROOM   WSType = "JOIN_ROOM"
	TEAM        WSType = "TEAM"
	GUESS       WSType = "GUESS"
	VALIDATE    WSType = "VALIDATE"
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
		slog.Error("Error upgrading connection", "error", err)
		return
	}

	client := &models.Client{
		User: models.User{
			Id:       generateUserId(),
			Username: "guest",
			Lives:    3,
		},
		Conn: conn,
		Send: make(chan []byte),
	}

	handleRead(client)
	go handleWrite(client)
}

// process incoming client messages
func handleRead(client *models.Client) {
	defer client.Conn.Close()

	var msg WSMessage
	for {
		err := client.Conn.ReadJSON(&msg)
		if err != nil {
			slog.Error("reading JSON", "error", err)
			return
		}
		switch {
		case msg.Type == CREATE_ROOM:
			handleCreateRoom(msg.Payload, client)
			slog.Info("room created", "client", client.User.Id)
		}
	}

}

func handleWrite(client *models.Client) {
	// send message to client/ broadcast to all clients
}

func generateUserId() string {
	return fmt.Sprintf("user_%d", time.Now().UnixNano())
}
