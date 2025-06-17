package game

import (
	"fmt"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
	"github.com/JaanLavaerts/ticktickbrick/internal/models"
	"github.com/JaanLavaerts/ticktickbrick/internal/room"
	"github.com/JaanLavaerts/ticktickbrick/internal/util"
)

func NextTurn(room *models.Room, teams []models.Team) {
	newTeam := data.RandomTeam(teams)
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

func SubmitGuess(r *models.Room, userId string, player models.Player) (bool, error) {
	answer := data.PlayerPlayedFor(player, r.CurrentTeam)
	client, ok := r.Clients[userId]

	if !ok {
		return false, fmt.Errorf(util.UserNotInRoomError)
	}

	if !room.AllUsersReady(r) {
		return false, fmt.Errorf("not all players are ready")
	}

	if !isClientTurn(r, userId) {
		return false, fmt.Errorf("not your turn")
	}

	if isPlayerMentioned(player, r.MentionedPlayers) {
		return false, fmt.Errorf("player already mentioned")
	}

	r.MentionedPlayers = append(r.MentionedPlayers, player)

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

func isClientTurn(room *models.Room, userId string) bool {
	return room.CurrentTurn < len(room.TurnOrder) && room.TurnOrder[room.CurrentTurn] == userId
}
