package main

import (
	"fmt"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
)

func main() {
	players, err := data.LoadData[data.Player]("assets/players.json")
	if err != nil {
		fmt.Println(err)
	}
	teams, err := data.LoadData[data.Team]("assets/teams.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(players, teams)
}
