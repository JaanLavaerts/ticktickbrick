package handlers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"

// 	"github.com/JaanLavaerts/ticktickbrick/internal/data"
// 	"github.com/JaanLavaerts/ticktickbrick/internal/models"
// 	"github.com/JaanLavaerts/ticktickbrick/internal/room"
// 	"github.com/JaanLavaerts/ticktickbrick/internal/util"
// )

// type APIResponse struct {
// 	StatusCode int    `json:"status_code"`
// 	Msg        string `json:"msg"`
// 	Data       any    `json:"data,omitempty"`
// }

// func writeResponse(w http.ResponseWriter, status int, msg string, data any) {
// 	response := &APIResponse{
// 		StatusCode: status,
// 		Msg:        msg,
// 		Data:       data,
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	if data != nil {
// 		log.Printf("%d - %s: %s", status, msg, data)
// 	} else {
// 		log.Printf("%d - %s", status, msg)
// 	}
// 	json.NewEncoder(w).Encode(response)
// }

// func Ping(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "pong")
// }

// func CreateRoomHandler(teams []models.Team) http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		team := data.RandomTeam(teams)

// 		body, err := io.ReadAll(req.Body)
// 		if err != nil {
// 			writeResponse(w, 400, util.InvalidInputError, nil)
// 			return
// 		}
// 		defer req.Body.Close()

// 		var client models.Client
// 		err = json.Unmarshal(body, &client)
// 		if err != nil {
// 			http.Error(w, err.Error(), 400)
// 			return
// 		}

// 		createdRoom, err := room.CreateRoom(&client, team)
// 		if err != nil {
// 			writeResponse(w, 409, util.UserAlreadyInRoomError, createdRoom.Id)
// 			return
// 		}

// 		room.Manager.AddRoom(&createdRoom)
// 		writeResponse(w, 200, util.RoomCreatedSuccess, createdRoom.Id)
// 	}

// }

// func JoinRoomHandler(w http.ResponseWriter, req *http.Request) {
// 	body, err := io.ReadAll(req.Body)
// 	if err != nil {
// 		writeResponse(w, 400, util.InvalidInputError, nil)
// 		return
// 	}
// 	defer req.Body.Close()

// 	var payload struct {
// 		RoomId string      `json:"room_id"`
// 		User   models.User `json:"user"`
// 	}

// 	err = json.Unmarshal(body, &payload)
// 	if err != nil {
// 		http.Error(w, err.Error(), 400)
// 		return
// 	}

// 	roomToJoin, err := room.Manager.GetRoom(payload.RoomId)
// 	if err != nil {
// 		writeResponse(w, 404, util.RoomNotFoundError, nil)
// 		return
// 	}
// 	err = room.JoinRoom(roomToJoin, payload.User)
// 	if err != nil {
// 		writeResponse(w, 400, util.UserAlreadyInRoomError, roomToJoin.Id)
// 		return
// 	}
// 	writeResponse(w, 200, util.UserJoinedRoomSuccess, roomToJoin.Id)
// }

// func GetAllRooms(w http.ResponseWriter, req *http.Request) {
// 	allRooms, err := room.Manager.GetAllRooms()
// 	if err != nil {
// 		writeResponse(w, 404, util.NoRoomsError, nil)
// 		return
// 	}
// 	writeResponse(w, 200, util.RoomsRetrievedSuccess, allRooms)
// }
