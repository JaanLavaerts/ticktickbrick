package handlers

import (
	"encoding/json"
	"log/slog"

	"github.com/JaanLavaerts/ticktickbrick/internal/models"
	"github.com/JaanLavaerts/ticktickbrick/internal/room"
)

type CreateRoomPayload struct {
	Username string `json:"username"`
}

type JoinRoomPayload struct {
	Username string `json:"username"`
	RoomId   string `json:"room_id"`
}

type UserReadyPayload struct {
	IsReady bool `json:"is_ready"`
}

func handleCreateRoom(payload json.RawMessage, client *models.Client) {
	var createRoomPayload CreateRoomPayload
	err := json.Unmarshal(payload, &createRoomPayload)
	if err != nil {
		slog.Error("creating room", "error", err)
		sendMessage(client, ERROR, err.Error())
		return
	}

	if createRoomPayload.Username != "" {
		client.User.Username = createRoomPayload.Username
	}

	createdRoom, err := room.CreateRoom(client)
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

	if joinRoomPayload.Username != "" {
		client.User.Username = joinRoomPayload.Username
	}

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

	// sendMessage(client, ROOM_JOINED, NewRoomDTO(roomToJoin))
	broadcastRoomUpdate(client.Room)
	slog.Info("room joined", "room", roomToJoin.Id)
}

func handleReady(payload json.RawMessage, client *models.Client) {
	var userReadyPayload UserReadyPayload
	err := json.Unmarshal(payload, &userReadyPayload)
	if err != nil {
		slog.Error("user ready", "error", err)
		sendMessage(client, ERROR, err.Error())
		return
	}

	client.User.IsReady = userReadyPayload.IsReady
	broadcastRoomUpdate(client.Room)
	slog.Info("user ready", "user", client.User.Username)

	if room.AllUsersReady(client.Room) {
		room.SetRoomState(client.Room, models.INPROGRESS)
		broadcastRoomUpdate(client.Room)
		slog.Info("room inprogress", "room", client.Room.Id)
	}
}
