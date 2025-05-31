package game

import (
	"testing"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
)

func newTestRoom() Room {
	users := []User{
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

	teams, _ := data.LoadData[data.Team]("../../assets/teams.json")
	team := data.RandomTeam(teams)

	return StartGame(users, team)
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
	player := data.Player{
		Id:        "2544",
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
	// TODO
}

func TestRemoveLifeWhenAnswerIsWrong(t *testing.T) {
	room := newTestRoom()
	player := data.Player{
		Id:        "2544",
		Name:      "LeBron James",
		Positions: []string{"Forward"},
		Teams:     []string{"CLE", "MIA", "CLE", "LAL"},
	}
	team := data.Team{
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
	player := data.Player{
		Id:        "2544",
		Name:      "LeBron James",
		Positions: []string{"Forward"},
		Teams:     []string{"CLE", "MIA", "CLE", "LAL"},
	}
	team := data.Team{
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
