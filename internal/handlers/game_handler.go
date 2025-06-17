package handlers

import (
	"encoding/json"
	"log/slog"

	"github.com/JaanLavaerts/ticktickbrick/internal/game"
	"github.com/JaanLavaerts/ticktickbrick/internal/models"
	"github.com/JaanLavaerts/ticktickbrick/internal/room"
)

type GuessPayload struct {
	Player models.Player `json:"player"`
}

func handleGuess(payload json.RawMessage, client *models.Client, teams []models.Team) {
	var guessPayload GuessPayload
	err := json.Unmarshal(payload, &guessPayload)
	if err != nil {
		slog.Error("guessing", "error", err.Error())
		sendMessage(client, ERROR, err.Error())
		return
	}

	result, err := game.SubmitGuess(client.Room, client.User.Id, guessPayload.Player)
	if err != nil {
		slog.Error("guessing", "error", err.Error())
		sendMessage(client, ERROR, err.Error())
		return
	}

	if game.IsGameOver(client.Room) {
		room.SetRoomState(client.Room, models.ENDED)

		winner, err := game.GetWinner(client.Room)
		if err != nil {
			sendMessage(client, ERROR, err.Error())
			slog.Error("guessing", "error", err)
		}

		broadcastMessage(client.Room, GAME_OVER, winner)
		slog.Info("game_over", "winner", winner)
	} else {
		game.NextTurn(client.Room, teams)
	}

	sendMessage(client, GUESS_RESULT, result)
	broadcastRoomUpdate(client.Room)
	slog.Info("guess submitted", "guess", guessPayload.Player.Name)
}
