package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"maps"
	"net/http"
	"strconv"

	"github.com/JaanLavaerts/ticktickbrick/internal/room"
	"github.com/JaanLavaerts/ticktickbrick/internal/util"
)

type APIResponse struct {
	StatusCode int    `json:"status_code"`
	Msg        string `json:"msg"`
	Data       any    `json:"data,omitempty"`
}

func writeResponse(w http.ResponseWriter, status int, msg string, data any) {
	response := &APIResponse{
		StatusCode: status,
		Msg:        msg,
		Data:       data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func GetAllRooms(w http.ResponseWriter, req *http.Request) {
	allRooms, err := room.Manager.GetAllRooms()
	if err != nil {
		writeResponse(w, 404, err.Error(), nil)
		slog.Error(strconv.Itoa(404), "error", err.Error())
		return
	}

	var roomIds []string
	for k := range maps.Keys(allRooms) {
		roomIds = append(roomIds, k)
	}
	writeResponse(w, 200, util.RoomsRetrievedSuccess, roomIds)
	slog.Error(strconv.Itoa(200), "rooms", roomIds)
}

func GetRoom(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		writeResponse(w, 400, util.InvalidInputError, nil)
		slog.Error(strconv.Itoa(400), "error", err.Error())
		return
	}
	defer req.Body.Close()

	var roomId string
	_ = json.Unmarshal(body, &roomId)
	room, err := room.Manager.GetRoom(roomId)

	if err != nil {
		writeResponse(w, 404, err.Error(), nil)
		slog.Error(strconv.Itoa(404), "error", err.Error())
		return
	}

	writeResponse(w, 200, util.RoomRetrievedSuccess, NewRoomDTO(room))
	slog.Error(strconv.Itoa(200), "room", NewRoomDTO(room))
}
