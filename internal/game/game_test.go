package game

import (
	"testing"
	"time"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
	"github.com/JaanLavaerts/ticktickbrick/internal/models"
)

func newTestRoom() models.Room {
	users := []models.User{
		{
			Id:       "1",
			Username: "userOne",
			Lives:    3,
		},
		{
			Id:       "2",
			Username: "userTwo",
			Lives:    3,
		},
		{
			Id:       "3",
			Username: "userThree",
			Lives:    3,
		},
	}

	teams, _ := data.LoadData[models.Team]("../../assets/teams.json")
	team := data.RandomTeam(teams)
	room := models.Room{
		Id:               "123",
		Users:            users,
		CurrentTurn:      0,
		CurrentTeam:      team,
		MentionedPlayers: nil,
		State:            models.RoomState(models.INPROGRESS),
		StartTime:        time.Now(),
	}
	return room
}

func TestNextTurnUserAlive(t *testing.T) {
	room := newTestRoom()
	// check if nextturn is on index 1
	nextTurn := room.CurrentTurn + 1
	NextTurn(&room, room.CurrentTeam)

	if room.CurrentTurn != nextTurn {
		t.Errorf("got %v, wanted %v", room.CurrentTurn, nextTurn)
	}
}

func TestNextTurnUserDead(t *testing.T) {
	room := newTestRoom()
	// check if nextturn is on index 2 if index 1 is dead
	nextTurn := room.CurrentTurn + 2

	room.Users[1].Lives = 0
	NextTurn(&room, room.CurrentTeam)

	if room.CurrentTurn != nextTurn {
		t.Errorf("got %v, wanted %v", room.CurrentTurn, nextTurn)
	}
}

func TestNextTurnUserAliveWrap(t *testing.T) {
	room := newTestRoom()
	// start at last player, check if it wraps to first player
	room.CurrentTurn = len(room.Users) - 1
	nextTurn := 0

	NextTurn(&room, room.CurrentTeam)

	if room.CurrentTurn != nextTurn {
		t.Errorf("got %v, wanted %v", room.CurrentTurn, nextTurn)
	}
}
func TestNextTurnUserDeadWrap(t *testing.T) {
	room := newTestRoom()
	// start at last player, check if it wraps to second player if first player is dead
	room.CurrentTurn = len(room.Users) - 1
	nextTurn := 1

	room.Users[0].Lives = 0
	NextTurn(&room, room.CurrentTeam)

	if room.CurrentTurn != nextTurn {
		t.Errorf("got %v, wanted %v", room.CurrentTurn, nextTurn)
	}
}

func TestSubmitAnswer(t *testing.T) {
	room := newTestRoom()
	player := models.Player{
		Id:        2544,
		Name:      "LeBron James",
		Positions: []string{"Forward"},
		Teams:     []string{"CLE", "MIA", "CLE", "LAL"},
	}

	SubmitAnswer(&room, "1", player)

	userHasAnswered := room.Users[0].HasAnswered == true
	playerIsMentioned := room.MentionedPlayers[0].Id == player.Id

	if !userHasAnswered {
		t.Errorf("got %v, wanted %v", userHasAnswered, true)
	}
	if !playerIsMentioned {
		t.Errorf("got %v, wanted %v", playerIsMentioned, true)
	}
}

func TestNotAbleToSubmitMentionedPlayer(t *testing.T) {
	room := newTestRoom()
	player := models.Player{
		Id:        2544,
		Name:      "LeBron James",
		Positions: []string{"Forward"},
		Teams:     []string{"CLE", "MIA", "CLE", "LAL"},
	}
	room.MentionedPlayers = append(room.MentionedPlayers, player)
	answer, err := SubmitAnswer(&room, "1", player)

	if answer != false || err == nil {
		t.Errorf("got %v, wanted %v", answer, err)
	}
}

func TestRemoveLifeWhenAnswerIsWrong(t *testing.T) {
	room := newTestRoom()
	player := models.Player{
		Id:        2544,
		Name:      "LeBron James",
		Positions: []string{"Forward"},
		Teams:     []string{"CLE", "MIA", "CLE", "LAL"},
	}
	team := models.Team{
		Name:         "Indiana Pacers",
		Abbreviation: "IND",
	}
	room.CurrentTeam = team
	SubmitAnswer(&room, "1", player)
	wasLifeRemoved := room.Users[0].Lives == 2

	if !wasLifeRemoved {
		t.Errorf("got %v, wanted %v", wasLifeRemoved, true)
	}
}

func TestNoLifeLostWhenAnswerIsRight(t *testing.T) {
	room := newTestRoom()
	player := models.Player{
		Id:        2544,
		Name:      "LeBron James",
		Positions: []string{"Forward"},
		Teams:     []string{"CLE", "MIA", "CLE", "LAL"},
	}
	team := models.Team{
		Name:         "Miami Heat",
		Abbreviation: "MIA",
	}
	room.CurrentTeam = team
	SubmitAnswer(&room, "1", player)
	wasNoLifeRemoved := room.Users[0].Lives == 3

	if !wasNoLifeRemoved {
		t.Errorf("got %v, wanted %v", wasNoLifeRemoved, true)
	}
}

func TestIsGameOver(t *testing.T) {
	room := newTestRoom()
	for i := range room.Users {
		room.Users[i].Lives = 0
	}

	isGameOver := IsGameOver(&room)
	if !isGameOver {
		t.Errorf("got %v, wanted %v", isGameOver, true)
	}

}

func TestGetWinnerGameIsOver(t *testing.T) {
	room := newTestRoom()

	for i := range room.Users {
		room.Users[i].Lives = 0
	}
	room.Users[2].Lives = 1 // winner stays
	winner, _ := GetWinner(&room)

	if winner != "userThree" {
		t.Errorf("got %v, wanted %v", winner, "userThree")
	}
}

func TestGetWinnerGameIsNotOver(t *testing.T) {
	room := newTestRoom()
	winner, err := GetWinner(&room)

	if winner != "" || err == nil {
		t.Errorf("got %v, wanted %v", winner, err)
	}
}
