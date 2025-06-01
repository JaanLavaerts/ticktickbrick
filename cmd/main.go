package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
	"github.com/JaanLavaerts/ticktickbrick/internal/handlers"
)

func main() {

	Players, err := data.LoadData[data.Player]("assets/players.json")
	if err != nil {
		fmt.Println(err)
	}
	Teams, err := data.LoadData[data.Team]("assets/teams.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Players, Teams)

	http.HandleFunc("/ping", handlers.Ping)
	http.HandleFunc("/create-room", handlers.CreateRoom)

	log.Println("server running on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
