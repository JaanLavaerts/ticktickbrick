package game

import (
	"fmt"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
	"github.com/JaanLavaerts/ticktickbrick/internal/models"
	"github.com/JaanLavaerts/ticktickbrick/internal/util"
)

func NextTurn(room *models.Room, newTeam models.Team) {
	room.CurrentTeam = newTeam
	startIndex := room.CurrentTurn

	for i := 1; i <= len(room.Clients); i++ {
		nextIndex := (startIndex + i) % len(room.Clients)
		nextClient := room.Clients[room.TurnOrder[nextIndex]]

		if nextClient.User.Lives > 0 {
			room.CurrentTurn = nextIndex
			return
		}
	}
}

func SubmitAnswer(room *models.Room, userId string, player models.Player) (bool, error) {
	answer := data.PlayerPlayedFor(player, room.CurrentTeam)
	client, ok := room.Clients[userId]

	if !ok {
		return false, fmt.Errorf(util.UserNotInRoomError)
	}

	if isPlayerMentioned(player, room.MentionedPlayers) {
		return false, fmt.Errorf("player already mentioned")
	}

	room.MentionedPlayers = append(room.MentionedPlayers, player)

	if !answer {
		client.User.Lives--
	}

	client.User.HasAnswered = true
	return answer, nil
}

func RemoveLife(room *models.Room, userId string) error {
	client, ok := room.Clients[userId]
	if !ok {
		return fmt.Errorf(util.UserNotInRoomError)
	}
	client.User.Lives--
	return nil
}

func IsGameOver(room *models.Room) bool {
	clients := getAliveClients(room)
	return len(clients) <= 1
}

func GetWinner(room *models.Room) (string, error) {
	clients := getAliveClients(room)
	if len(clients) != 1 {
		return "", fmt.Errorf(util.GameNotOverYetError)
	}

	return clients[0].User.Username, nil
}

func getAliveClients(room *models.Room) []*models.Client {
	var alive []*models.Client
	for _, client := range room.Clients {
		if client.User.Lives > 0 {
			alive = append(alive, client)
		}
	}
	return alive
}

func isPlayerMentioned(player models.Player, players []models.Player) bool {
	for i := range players {
		if players[i].Id == player.Id {
			return true
		}
	}
	return false
}
