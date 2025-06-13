package handlers

import (
	"encoding/json"
	"log/slog"

	"github.com/JaanLavaerts/ticktickbrick/internal/models"
	"github.com/JaanLavaerts/ticktickbrick/internal/room"
)

type CreateRoomPayload struct {
	User models.User `json:"user"`
}

func handleCreateRoom(payload json.RawMessage, client *models.Client, team models.Team) error {
	var createRoomPayload CreateRoomPayload
	err := json.Unmarshal(payload, &createRoomPayload)
	if err != nil {
		slog.Error("creating room", "error", err)
		sendMessage(client, ERROR, err.Error())
		return err
	}
	client.User = createRoomPayload.User

	createdRoom, err := room.CreateRoom(client, team)
	if err != nil {
		slog.Error("creating room", "error", err)
		sendMessage(client, ERROR, err.Error())
		return err
	}

	sendMessage(client, ROOM_CREATED, createdRoom.Id)
	slog.Info("room created", "room", createdRoom.Id)
	return nil
}
