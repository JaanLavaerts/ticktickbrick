package game

import (
	"fmt"
	"time"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
	"github.com/JaanLavaerts/ticktickbrick/internal/models"
)

func StartGame(users []models.User, team models.Team) models.Room {
	randomId := generateTimestampID()
	room := &models.Room{
		Id:               randomId,
		Users:            users,
		CurrentTurn:      0, // TODO: should be random in the future, be ware of NextTurn logic/ tests
		CurrentTeam:      team,
		MentionedPlayers: nil,
		State:            models.RoomState(models.INPROGRESS),
		StartTime:        time.Now(),
	}
	return *room
}

func NextTurn(room *models.Room, newTeam models.Team) {
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

func SubmitAnswer(room *models.Room, userId string, player models.Player) (bool, error) {
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

func RemoveLife(room *models.Room, userId string) {
	userIdx := getUserIdxById(room.Users, userId)
	room.Users[userIdx].Lives--
}

func IsGameOver(room *models.Room) bool {
	var aliveCount int

	for i := range room.Users {
		user := room.Users[i]
		if user.Lives != 0 {
			aliveCount++
		}
	}
	return aliveCount <= 1
}

func GetWinner(room *models.Room) (string, error) {
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
func getUserIdxById(users []models.User, userId string) int {
	var idx int
	for i := range users {
		if users[i].Id == userId {
			idx = i
		}
	}
	return idx
}

func isPlayerMentioned(player models.Player, players []models.Player) bool {
	for i := range players {
		if players[i].Id == player.Id {
			return true
		}
	}
	return false
}

func generateTimestampID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
