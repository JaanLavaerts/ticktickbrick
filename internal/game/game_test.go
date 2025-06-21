package game

import (
	"testing"
	"time"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
	"github.com/JaanLavaerts/ticktickbrick/internal/models"
)

func newTestRoom() *models.Room {
	err := data.LoadTeams("../../assets/teams.json")
	if err != nil {
		panic("error loading teams: " + err.Error())
	}
	team := data.RandomTeam()

	room := &models.Room{
		Id:               "123",
		Clients:          make(map[string]*models.Client),
		CurrentTurn:      0,
		TurnOrder:        []string{"1", "2", "3"},
		CurrentTeam:      team,
		MentionedPlayers: nil,
		State:            models.RoomState(models.INPROGRESS),
		StartTime:        time.Now(),
	}

	room.Clients["1"] = &models.Client{
		User: models.User{Id: "1", Username: "user_1", Lives: 3, IsReady: true},
	}
	room.Clients["2"] = &models.Client{
		User: models.User{Id: "2", Username: "user_2", Lives: 3, IsReady: true},
	}
	room.Clients["3"] = &models.Client{
		User: models.User{Id: "3", Username: "user_3", Lives: 3, IsReady: true},
	}

	return room
}

func TestNextTurnNewTeam(t *testing.T) {
	room := newTestRoom()
	previousTeam := room.CurrentTeam
	NextTurn(room)
	if room.CurrentTeam == previousTeam {
		t.Errorf("got %v, wanted", room.CurrentTeam)
	}
}

func TestNextTurnUserAlive(t *testing.T) {
	room := newTestRoom()
	// check if nextturn is on index 1
	nextTurn := room.CurrentTurn + 1
	NextTurn(room)

	if room.CurrentTurn != nextTurn {
		t.Errorf("got %v, wanted %v", room.CurrentTurn, nextTurn)
	}
}

func TestNextTurnUserDead(t *testing.T) {
	room := newTestRoom()
	// check if nextturn is on index 2 if index 1 is dead
	nextTurn := room.CurrentTurn + 2

	room.Clients["2"].User.Lives = 0
	NextTurn(room)

	if room.CurrentTurn != nextTurn {
		t.Errorf("got %v, wanted %v", room.CurrentTurn, nextTurn)
	}
}

func TestNextTurnUserAliveWrap(t *testing.T) {
	room := newTestRoom()
	// start at last player, check if it wraps to first player
	room.CurrentTurn = len(room.Clients) - 1
	nextTurn := 0

	NextTurn(room)

	if room.CurrentTurn != nextTurn {
		t.Errorf("got %v, wanted %v", room.CurrentTurn, nextTurn)
	}
}

func TestNextTurnUserDeadWrap(t *testing.T) {
	room := newTestRoom()
	// start at last player, check if it wraps to second player if first player is dead
	room.CurrentTurn = len(room.Clients) - 1
	nextTurn := 1

	room.Clients["1"].User.Lives = 0
	NextTurn(room)

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

	SubmitGuess(room, "1", player)

	userHasAnswered := room.Clients["1"].User.HasAnswered == true
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
	answer, err := SubmitGuess(room, "1", player)

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
	SubmitGuess(room, "1", player)
	wasLifeRemoved := room.Clients["1"].User.Lives == 2

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
	SubmitGuess(room, "1", player)
	wasNoLifeRemoved := room.Clients["1"].User.Lives == 3

	if !wasNoLifeRemoved {
		t.Errorf("got %v, wanted %v", wasNoLifeRemoved, true)
	}
}

func TestIsGameOver(t *testing.T) {
	room := newTestRoom()
	for i := range room.Clients {
		room.Clients[i].User.Lives = 0
	}

	isGameOver := IsGameOver(room)
	if !isGameOver {
		t.Errorf("got %v, wanted %v", isGameOver, true)
	}

}

func TestGetWinnerGameIsOver(t *testing.T) {
	room := newTestRoom()

	for i := range room.Clients {
		room.Clients[i].User.Lives = 0
	}
	room.Clients["3"].User.Lives = 1 // winner stays
	winner, _ := GetWinner(room)

	if winner != "user_3" {
		t.Errorf("got %v, wanted %v", winner, "user_3")
	}
}

func TestGetWinnerGameIsNotOver(t *testing.T) {
	room := newTestRoom()
	winner, err := GetWinner(room)

	if winner != "" || err == nil {
		t.Errorf("got %v, wanted %v", winner, err)
	}
}
