package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
	"github.com/JaanLavaerts/ticktickbrick/internal/models"
	"github.com/JaanLavaerts/ticktickbrick/internal/room"
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

func Ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong")
}

func CreateRoomHandler(teams []models.Team) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		team := data.RandomTeam(teams)

		body, err := io.ReadAll(req.Body)
		if err != nil {
			msg := "invalid body"
			writeResponse(w, 400, msg, nil)
			return
		}
		defer req.Body.Close()

		var user models.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		createdRoom, err := room.CreateRoom(user, team)
		if err != nil {
			msg := "you already have a room"
			writeResponse(w, 409, msg, nil)
			log.Printf("%d user already has a room: %s", 409, createdRoom.Id)
			return
		}

		room.Manager.AddRoom(&createdRoom)
		msg := "room created succesfully"
		log.Printf("%d new room created: %s", 200, createdRoom.Id)
		writeResponse(w, 200, msg, nil)
	}

}

func JoinRoom(w http.ResponseWriter, req *http.Request) {

}
