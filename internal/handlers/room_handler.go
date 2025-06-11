package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/JaanLavaerts/ticktickbrick/internal/models"
)

func handleCreateRoom(payload json.RawMessage, client *models.Client) {
	fmt.Println("room created")
}
