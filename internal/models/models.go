package models

import (
	"time"

	"github.com/gorilla/websocket"
)

// room types
type RoomState int

const (
	WAITING    RoomState = iota // auto increment = 0
	INPROGRESS                  // = 1
	ENDED                       // = 2
)

type Room struct {
	Id               string             `json:"id"`
	Clients          map[string]*Client `json:"clients"`
	CurrentTurn      int                `json:"current_turn"` // index of a user in []Users
	TurnOrder        []string           `json:"turn_order"`   // slice of client IDs
	MentionedPlayers []Player           `json:"mentioned_players"`
	CurrentTeam      Team               `json:"current_team"`
	State            RoomState          `json:"state"`
	StartTime        time.Time          `json:"start_time"`
}

type Client struct {
	User User `json:"user"`
	Conn *websocket.Conn
	Room *Room
	Send chan []byte
}

// game types
type User struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Lives       int    `json:"lives"`
	HasAnswered bool   `json:"has_answered"`
	IsReady     bool   `json:"is_ready"`
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
