package game

import (
	"time"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
)

func StartGame(users []User, team data.Team) Room {
	room := &Room{
		Users:            users,
		CurrentTurn:      0, // TODO: should be random in the future, be ware of NextTurn logic
		CurrentTeam:      team,
		MentionedPlayers: nil,
		State:            RoomState(INPROGRESS),
		StartTime:        time.Now(),
	}
	return *room
}

func NextTurn(room *Room, teams []data.Team) {
	newTeam := data.RandomTeam(teams)
	room.CurrentTeam = newTeam

	isNextUserAlive := room.Users[room.CurrentTurn+1].Lives != 0

	if room.CurrentTurn == len(room.Users)-1 && isNextUserAlive {
		room.CurrentTurn = 0
	} else if isNextUserAlive {
		room.CurrentTurn = room.CurrentTurn + 1
	} else {
		room.State = RoomState(ENDED)
	}
}
