package game

import (
	"testing"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
)

var teams, _ = data.LoadData[data.Team]("../../assets/teams.json")
var team = data.RandomTeam(teams)

var userOne = User{
	Id:          "1",
	Username:    "userOne",
	Lives:       3,
	HasAnswered: false,
}
var userTwo = User{
	Id:          "2",
	Username:    "userTwo",
	Lives:       3,
	HasAnswered: false,
}
var userThree = User{
	Id:          "3",
	Username:    "userThree",
	Lives:       3,
	HasAnswered: false,
}

var users = []User{userOne, userTwo, userThree}
var room = StartGame(users, team)

func TestNextTurnUserAlive(t *testing.T) {
	// check if nextturn is on index 1
	nextTurn := room.CurrentTurn + 1
	NextTurn(&room, team)

	if room.CurrentTurn != nextTurn {
		t.Errorf("got %v, wanted %v", room.CurrentTurn, nextTurn)
	}
}

func TestNextTurnUserDead(t *testing.T) {
	// check if nextturn is on index 2 if index 1 is dead
	nextTurn := room.CurrentTurn + 2

	room.Users[1].Lives = 0
	NextTurn(&room, team)

	if room.CurrentTurn != nextTurn {
		t.Errorf("got %v, wanted %v", room.CurrentTurn, nextTurn)
	}
}

func TestNextTurnUserAliveWrap(t *testing.T) {
	// start at last player, check if it wraps to first player
	room.CurrentTurn = len(room.Users) - 1
	nextTurn := 0

	NextTurn(&room, team)

	if room.CurrentTurn != nextTurn {
		t.Errorf("got %v, wanted %v", room.CurrentTurn, nextTurn)
	}
}
func TestNextTurnUserDeadWrap(t *testing.T) {
	// start at last player, check if it wraps to second player if first player is dead
	room.CurrentTurn = len(room.Users) - 1
	nextTurn := 1

	room.Users[0].Lives = 0
	NextTurn(&room, team)

	if room.CurrentTurn != nextTurn {
		t.Errorf("got %v, wanted %v", room.CurrentTurn, nextTurn)
	}
}
