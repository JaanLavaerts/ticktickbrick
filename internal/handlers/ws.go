package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"maps"
	"net/http"
	"time"

	"github.com/JaanLavaerts/ticktickbrick/internal/models"
	"github.com/gorilla/websocket"
)

type WSType string

const (
	CREATE_ROOM  WSType = "CREATE_ROOM"
	ROOM_CREATED WSType = "ROOM_CREATED"
	UPDATE_ROOM  WSType = "UPDATE_ROOM"
	JOIN_ROOM    WSType = "JOIN_ROOM"
	USER_GUESS   WSType = "USER_GUESS"
	GUESS_RESULT WSType = "GUESS_RESULT"
	GAME_OVER    WSType = "GAME_OVER"
	USER_READY   WSType = "USER_READY"
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
		handleRead(client, teams)
	}
}

// process incoming client messages
func handleRead(client *models.Client, teams []models.Team) {
	defer client.Conn.Close()

	var msg WSMessage
	for {
		err := client.Conn.ReadJSON(&msg)
		if err != nil {
			slog.Error("reading JSON", "error", err)
			return
		}
		switch msg.Type {
		case CREATE_ROOM:
			handleCreateRoom(msg.Payload, client, teams)
		case JOIN_ROOM:
			handleJoinRoom(msg.Payload, client)
		case USER_READY:
			handleReady(msg.Payload, client)
		case USER_GUESS:
			handleGuess(msg.Payload, client, teams)
		default:
			slog.Error("not a supported WSType", "", nil)
		}
	}

}

// send message to client
func handleWrite(client *models.Client) {
	for msg := range client.Send {
		err := client.Conn.WriteMessage(websocket.TextMessage, msg)

		if err != nil {
			slog.Error("write", "error", err)
			break
		}
	}
}

func broadcastRoomUpdate(room *models.Room) error {
	newRoom := NewRoomDTO(room)

	for client := range maps.Values(room.Clients) {
		sendMessage(client, UPDATE_ROOM, newRoom)
	}
	return nil
}

func broadcastMessage(room *models.Room, messageType WSType, rawPayload any) error {
	for client := range maps.Values(room.Clients) {
		sendMessage(client, messageType, rawPayload)
	}
	return nil
}

func sendMessage(client *models.Client, messageType WSType, rawPayload any) error {
	payload, err := json.Marshal(rawPayload)
	if err != nil {
		return err
	}

	response := &WSMessage{
		Type:    messageType,
		Payload: payload,
	}

	data, err := json.Marshal(response)
	if err != nil {
		return err
	}

	client.Send <- data
	return nil
}

func generateUserId() string {
	return fmt.Sprintf("user_%d", time.Now().UnixNano())
}
