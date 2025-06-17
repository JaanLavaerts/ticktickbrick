package handlers

import (
	"time"

	"github.com/JaanLavaerts/ticktickbrick/internal/models"
)

// DTO with users instead of clients to not expose server stuff like conn, send
type RoomDTO struct {
	Id               string           `json:"id"`
	Users            []models.User    `json:"users"`
	CurrentTurn      int              `json:"current_turn"`
	TurnOrder        []string         `json:"turn_order"`
	MentionedPlayers []models.Player  `json:"mentioned_players"`
	CurrentTeam      models.Team      `json:"current_team"`
	State            models.RoomState `json:"state"`
	StartTime        time.Time        `json:"start_time"`
}

func NewRoomDTO(room *models.Room) RoomDTO {
	var users []models.User
	for _, i := range room.Clients {
		users = append(users, i.User)
	}
	return RoomDTO{
		Id:               room.Id,
		Users:            users,
		CurrentTurn:      room.CurrentTurn,
		TurnOrder:        room.TurnOrder,
		MentionedPlayers: room.MentionedPlayers,
		CurrentTeam:      room.CurrentTeam,
		State:            room.State,
		StartTime:        room.StartTime,
	}
}
