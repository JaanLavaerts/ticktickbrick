package data

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"slices"

	"github.com/JaanLavaerts/ticktickbrick/internal/models"
)

var Teams []models.Team
var Players []models.Player

func LoadData[T any](fileString string) ([]T, error) {
	var data []T

	file, err := os.Open(fileString)
	if err != nil {
		return data, fmt.Errorf("error while opening file: %w", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return data, fmt.Errorf("error while reading file: %w", err)
	}

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return data, fmt.Errorf("error while unmarshaling file: %w", err)
	}

	return data, nil
}

func LoadTeams(fileString string) error {
	teams, err := LoadData[models.Team](fileString)
	if err != nil {
		return fmt.Errorf("error loading teams: %w", err)
	}
	Teams = teams
	return nil
}

func LoadPlayers(fileString string) error {
	players, err := LoadData[models.Player](fileString)
	if err != nil {
		return fmt.Errorf("error loading players: %w", err)
	}
	Players = players
	return nil
}

func PlayerPlayedFor(player models.Player, team models.Team) bool {
	return slices.Contains(player.Teams, team.Abbreviation)
}

func RandomTeam() models.Team {
	return Teams[rand.IntN(len(Teams))]
}
