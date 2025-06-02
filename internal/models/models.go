package models

import (
	"time"
)

// room types
type RoomState int

const (
	WAITING    RoomState = iota // auto increment = 0
	INPROGRESS                  // = 1
	ENDED                       // = 2
)

type Room struct {
	Id               string
	Users            []User
	CurrentTurn      int // index of a user in []Users
	MentionedPlayers []Player
	CurrentTeam      Team
	State            RoomState
	StartTime        time.Time
}

// game types
type User struct {
	Id          string
	Username    string
	Lives       int
	HasAnswered bool // has the player answered yet this round
}

// data types
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
