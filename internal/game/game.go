package game

import (
	"time"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
)

type RoomState int

const (
	WAITING    RoomState = iota // auto increment = 0
	INPROGRESS                  // = 1
	ENDED                       // = 2
)

type Room struct {
	Users            []User
	CurrentTurn      int // index of a user in []Users
	MentionedPlayers []data.Player
	CurrentTeam      data.Team
	State            RoomState
	StartTime        time.Time
}

type User struct {
	Id          string
	Username    string
	Lives       int
	HasAnswered bool // has the player answered yet this round
}

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

func SubmitAnswer(room *Room, userId string, player data.Player) (bool, error) {
	answer := data.PlayerPlayedFor(player, room.CurrentTeam)
	room.MentionedPlayers = append(room.MentionedPlayers, player)

	//TODO: remove life of user when answer is wrong
	//TODO: set hasAnswered to true

	return answer, nil
}

func RemoveLife(room *Room, userId string) {
	user := getUserById(room.Users, userId)
	user.Lives--
}

func IsGameOver(room *Room) bool {
	var aliveCount int

	for i := range room.Users {
		user := room.Users[i]
		if user.Lives != 0 {
			aliveCount++
		}
	}
	return aliveCount <= 1
}

func GetWinner(room *Room) string {
	var winnerUsername string

	for i := range room.Users {
		user := room.Users[i]
		if user.Lives != 0 {
			winnerUsername = user.Username
		}
	}
	return winnerUsername
}

func getUserById(users []User, userId string) User {
	var user User
	for i := range users {
		if users[i].Id == userId {
			user = users[i]
		}
	}
	return user
}
