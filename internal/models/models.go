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
	Id               string    `json:"id"`
	Users            []User    `json:"users"`
	CurrentTurn      int       `json:"current_turn"` // index of a user in []Users
	MentionedPlayers []Player  `json:"mentioned_players"`
	CurrentTeam      Team      `json:"current_team"`
	State            RoomState `json:"state"`
	StartTime        time.Time `json:"start_time"`
}

// game types
type User struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Lives       int    `json:"lives"`
	HasAnswered bool   `json:"has_answered"`
}

// data types
type Team struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

type Player struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Positions []string `json:"positions"`
	Teams     []string `json:"teams"`
}
