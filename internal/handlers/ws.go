package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
	"github.com/JaanLavaerts/ticktickbrick/internal/models"
	"github.com/gorilla/websocket"
)

type WSType string

const (
	CREATE_ROOM  WSType = "CREATE_ROOM"
	ROOM_CREATED WSType = "ROOM_CREATED"
	JOIN_ROOM    WSType = "JOIN_ROOM"
	TEAM         WSType = "TEAM"
	GUESS        WSType = "GUESS"
	VALIDATE     WSType = "VALIDATE"
	ERROR        WSType = "ERROR"
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

func WsHandler(teams []models.Team) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		team := data.RandomTeam(teams)

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			slog.Error("error upgrading connection", "error", err)
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

		go handleWrite(client)
		handleRead(client, team)
	}
}

// process incoming client messages
func handleRead(client *models.Client, team models.Team) {
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
			handleCreateRoom(msg.Payload, client, team)
		default:
			slog.Error("not a supported type", "", nil)
		}
	}

}

// send message to client/ broadcast to all clients
func handleWrite(client *models.Client) {
	for msg := range client.Send {
		err := client.Conn.WriteMessage(websocket.TextMessage, msg)

		if err != nil {
			slog.Error("write", "error", err)
			break
		}
	}
}

func sendMessage(client *models.Client, messageType WSType, rawPayload any) {
	payload, err := json.Marshal(rawPayload)
	if err != nil {
		slog.Error("marshaling payload", "error", err)
		return
	}

	response := &WSMessage{
		Type:    messageType,
		Payload: payload,
	}

	data, err := json.Marshal(response)
	if err != nil {
		slog.Error("marshaling response", "error", err)
		return
	}

	client.Send <- data
}

func generateUserId() string {
	return fmt.Sprintf("user_%d", time.Now().UnixNano())
}
