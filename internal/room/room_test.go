package room

import (
	"testing"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
	"github.com/JaanLavaerts/ticktickbrick/internal/game"
)

func newTestRoom() game.Room {
	users := []game.User{
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

	// reset global state
	roomManager.rooms = make(map[string]*game.Room)
	return game.StartGame(users, team)
}

func TestAddRoom(t *testing.T) {
	room := newTestRoom()
	roomManager.AddRoom(&room)

	if len(roomManager.rooms) <= 0 {
		t.Errorf("got %v, wanted %v", len(roomManager.rooms), 1)
	}
}

func TestGetRoom(t *testing.T) {
	room := newTestRoom()
	room.Id = "1"
	roomManager.AddRoom(&room)
	firstRoom := roomManager.GetRoom(room.Id)

	if firstRoom.Id != "1" {
		t.Errorf("got %v, wanted %v", len(roomManager.rooms), 1)
	}
}

func TestGetAllRooms(t *testing.T) {
	room := newTestRoom()
	roomManager.AddRoom(&room)
	allRoomsCount := len(roomManager.GetAllRooms())

	if allRoomsCount <= 0 {
		t.Errorf("got %v, wanted %v", len(roomManager.rooms), 1)
	}
}
