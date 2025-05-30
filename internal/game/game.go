package game

import (
	"fmt"
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

func NextTurn(room *Room, newTeam data.Team) {
	room.CurrentTeam = newTeam
	startIndex := room.CurrentTurn

	for i := 1; i <= len(room.Users); i++ {
		nextIndex := (startIndex + i) % len(room.Users)
		if room.Users[nextIndex].Lives > 0 {
			room.CurrentTurn = nextIndex
			return
		}
	}
}

func SubmitAnswer(room *Room, userId string, player data.Player) (bool, error) {
	userIdx := getUserIdxById(room.Users, userId)
	answer := data.PlayerPlayedFor(player, room.CurrentTeam)

	if isPlayerMentioned(player, room.MentionedPlayers) {
		return false, fmt.Errorf("player already mentioned")
	}

	room.MentionedPlayers = append(room.MentionedPlayers, player)

	if !answer {
		RemoveLife(room, userId)
	}

	room.Users[userIdx].HasAnswered = true
	return answer, nil
}

func RemoveLife(room *Room, userId string) {
	userIdx := getUserIdxById(room.Users, userId)
	room.Users[userIdx].Lives--
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

func GetWinner(room *Room) (string, error) {
	var winnerUsername string
	isGameOver := IsGameOver(room)

	if !isGameOver {
		return "", fmt.Errorf("game is not over yet, no winner")
	}

	for i := range room.Users {
		user := room.Users[i]
		if user.Lives != 0 {
			winnerUsername = user.Username
		}
	}
	return winnerUsername, nil
}

// helper functions
func getUserIdxById(users []User, userId string) int {
	var idx int
	for i := range users {
		if users[i].Id == userId {
			idx = i
		}
	}
	return idx
}

func isPlayerMentioned(player data.Player, players []data.Player) bool {
	for i := range players {
		if players[i].Id == player.Id {
			return true
		}
	}
	return false
}
