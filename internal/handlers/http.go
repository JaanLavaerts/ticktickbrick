package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "github.com/JaanLavaerts/ticktickbrick/internal/game"
	// "github.com/JaanLavaerts/ticktickbrick/internal/room"
)

func Ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong")
}

func CreateRoom(w http.ResponseWriter, req *http.Request) {

	body, err := io.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var id string
	err = json.Unmarshal(body, &id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(id)

}

func JoinRoom(w http.ResponseWriter, req *http.Request) {

}
