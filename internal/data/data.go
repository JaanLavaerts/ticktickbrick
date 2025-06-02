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

func PlayerPlayedFor(player models.Player, team models.Team) bool {
	return slices.Contains(player.Teams, team.Abbreviation)
}

func RandomTeam(teams []models.Team) models.Team {
	return teams[rand.IntN(len(teams))]
}
