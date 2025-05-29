package data

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Team struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

type Player struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Positions []string `json:"positions"`
	Teams     []string `json:"teams"`
}

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
