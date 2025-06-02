package room

import (
	"testing"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
	"github.com/JaanLavaerts/ticktickbrick/internal/game"
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

	// reset global state
	Manager.rooms = make(map[string]*models.Room)
	return game.StartGame(users, team)
}

func TestAddRoom(t *testing.T) {
	room := newTestRoom()
	Manager.AddRoom(&room)

	if len(Manager.rooms) <= 0 {
		t.Errorf("got %v, wanted %v", len(Manager.rooms), 1)
	}
}

func TestGetRoom(t *testing.T) {
	room := newTestRoom()
	room.Id = "1"
	Manager.AddRoom(&room)
	firstRoom := Manager.GetRoom(room.Id)

	if firstRoom.Id != "1" {
		t.Errorf("got %v, wanted %v", len(Manager.rooms), 1)
	}
}

func TestGetAllRooms(t *testing.T) {
	room := newTestRoom()
	Manager.AddRoom(&room)
	allRoomsCount := len(Manager.GetAllRooms())

	if allRoomsCount <= 0 {
		t.Errorf("got %v, wanted %v", len(Manager.rooms), 1)
	}
}

func TestJoinRoom(t *testing.T) {
	room := newTestRoom()
	Manager.AddRoom(&room)
	user := models.User{Id: "3", Username: "userThree", Lives: 3}
	JoinRoom(&room, user)

	if len(room.Users) != 4 {
		t.Errorf("got %v, wanted %v", len(room.Users), 4)
	}

}
