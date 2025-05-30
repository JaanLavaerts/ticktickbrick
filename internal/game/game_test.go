package game

import (
	"testing"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
)

var teams, _ = data.LoadData[data.Team]("assets/teams.json")
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

func TestNextTurnTeamChanged(t *testing.T) {
	// check new team has changed (could be same)
}
func TestNextTurnUserAlive(t *testing.T) {
	// check if nextturn is on index 1
}
func TestNextTurnUserDead(t *testing.T) {
	// check if nextturn is on index 2 if index 1 is dead
}
