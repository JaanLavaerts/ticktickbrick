package game

import "time"

type RoomState int

const (
	WAITING    RoomState = iota // auto increment = 0
	INPROGRESS                  // = 1
	ENDED                       // = 2
)

type Room struct {
	Users            []User
	CurrentTurn      string // id of a user
	MentionedPlayers []string
	State            RoomState
	StartTime        time.Time
}
