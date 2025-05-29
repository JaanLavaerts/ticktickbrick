package data

import "testing"

// stupid unit test to know syntax
func TestPlayerPlayedFor(t *testing.T) {
	player := &Player{
		Id:        "2544",
		Name:      "LeBron James",
		Positions: []string{"Forward"},
		Teams:     []string{"CLE", "MIA", "CLE", "LAL"},
	}

	team := &Team{
		Name:         "Miami Heat",
		Abbreviation: "MIA",
	}

	got := PlayerPlayedFor(*player, *team)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
