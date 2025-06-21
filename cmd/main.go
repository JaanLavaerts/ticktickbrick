package main

import (
	"log"
	"net/http"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
	"github.com/JaanLavaerts/ticktickbrick/internal/handlers"
)

func main() {

	err := data.LoadTeams("assets/teams.json")
	if err != nil {
		log.Fatalf("error loading teams: %v", err)
	}
	err = data.LoadPlayers("assets/players.json")
	if err != nil {
		log.Fatalf("error loading players: %v", err)
	}

	http.HandleFunc("/ws", handlers.WsHandler)

	// http routes
	http.HandleFunc("/rooms", handlers.GetAllRooms)
	http.HandleFunc("/room", handlers.GetRoom)

	log.Println("server running on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
