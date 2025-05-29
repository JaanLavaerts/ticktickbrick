package game

import (
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
	MentionedPlayers []string
	CurrentTeam      data.Team
	State            RoomState
	StartTime        time.Time
}
