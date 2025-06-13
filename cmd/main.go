package main

import (
	"log"
	"net/http"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
	"github.com/JaanLavaerts/ticktickbrick/internal/handlers"
	"github.com/JaanLavaerts/ticktickbrick/internal/models"
)

func main() {
	players, err := data.LoadData[models.Player]("assets/players.json")
	if err != nil {
		log.Fatal(players, err)
	}
	teams, err := data.LoadData[models.Team]("assets/teams.json")
	if err != nil {
		log.Fatal(teams, err)
	}

	// dependency injection of teams
	http.HandleFunc("/ws", handlers.WsHandler(teams))

	// http routes
	http.HandleFunc("/rooms", handlers.GetAllRooms)
	http.HandleFunc("/room", handlers.GetRoom)

	log.Println("server running on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
