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

type JoinRoomPayload struct {
	User   models.User `json:"user"`
	RoomId string      `json:"room_id"`
}

func handleCreateRoom(payload json.RawMessage, client *models.Client, team models.Team) {
	var createRoomPayload CreateRoomPayload
	err := json.Unmarshal(payload, &createRoomPayload)
	if err != nil {
		slog.Error("creating room", "error", err)
		sendMessage(client, ERROR, err.Error())
		return
	}
	client.User = createRoomPayload.User

	createdRoom, err := room.CreateRoom(client, team)
	if err != nil {
		slog.Error("creating room", "error", err)
		sendMessage(client, ERROR, err.Error())
		return
	}

	sendMessage(client, ROOM_CREATED, createdRoom.Id)
	slog.Info("room created", "room", createdRoom.Id)
}

func handleJoinRoom(payload json.RawMessage, client *models.Client) {
	var joinRoomPayload JoinRoomPayload
	err := json.Unmarshal(payload, &joinRoomPayload)
	if err != nil {
		slog.Error("joining room", "error", err)
		sendMessage(client, ERROR, err.Error())
		return
	}
	client.User = joinRoomPayload.User

	roomToJoin, err := room.Manager.GetRoom(joinRoomPayload.RoomId)
	if err != nil {
		slog.Error("joining room", "error", err)
		sendMessage(client, ERROR, err.Error())
		return
	}

	err = room.JoinRoom(roomToJoin, client)
	if err != nil {
		slog.Error("joining room", "error", err)
		sendMessage(client, ERROR, err.Error())
		return
	}

	sendMessage(client, ROOM_JOINED, NewRoomDTO(roomToJoin))
	slog.Info("room joined", "room", roomToJoin.Id)
}
