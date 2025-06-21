package room

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
		User: models.User{Id: "1", Username: "user_1", Lives: 3},
	}
	room.Clients["2"] = &models.Client{
		User: models.User{Id: "2", Username: "user_2", Lives: 3},
	}
	room.Clients["3"] = &models.Client{
		User: models.User{Id: "3", Username: "user_3", Lives: 3},
	}

	return room
}

func TestAddRoom(t *testing.T) {
	room := newTestRoom()
	Manager.AddRoom(room)

	if len(Manager.rooms) <= 0 {
		t.Errorf("got %v, wanted %v", len(Manager.rooms), 1)
	}
}

func TestGetRoom(t *testing.T) {
	room := newTestRoom()
	room.Id = "1"
	Manager.AddRoom(room)
	firstRoom, _ := Manager.GetRoom(room.Id)

	if firstRoom.Id != "1" {
		t.Errorf("got %v, wanted %v", len(Manager.rooms), 1)
	}
}

func TestGetAllRooms(t *testing.T) {
	room := newTestRoom()
	Manager.AddRoom(room)
	allRooms, _ := Manager.GetAllRooms()
	allRoomsCount := len(allRooms)

	if allRoomsCount <= 0 {
		t.Errorf("got %v, wanted %v", allRoomsCount, 1)
	}
}

func TestJoinRoom(t *testing.T) {
	room := newTestRoom()
	Manager.AddRoom(room)
	client := &models.Client{User: models.User{Id: "4", Username: "user_4", Lives: 3}, Conn: nil, Room: room, Send: nil}
	JoinRoom(room, client)

	if len(room.Clients) != 4 {
		t.Errorf("got %v, wanted %v", len(room.Clients), 4)
	}
}
