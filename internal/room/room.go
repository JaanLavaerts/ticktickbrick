package room

import (
	"fmt"
	"maps"

	"github.com/JaanLavaerts/ticktickbrick/internal/models"
)

type RoomManager struct {
	rooms map[string]*models.Room
}

var Manager = &RoomManager{
	rooms: make(map[string]*models.Room),
}

func CreateRoom(users []models.User, team models.Team) models.Room {
	return models.Room{}
}

func (r *RoomManager) AddRoom(room *models.Room) {
	r.rooms[room.Id] = room
}

func (r *RoomManager) GetRoom(id string) *models.Room {
	return r.rooms[id]
}

func (r *RoomManager) GetAllRooms() map[string]*models.Room {
	rooms := make(map[string]*models.Room)
	maps.Copy(rooms, r.rooms)
	return rooms
}

func (r *RoomManager) DoesRoomExist(room *models.Room) bool {
	_, ok := r.rooms[room.Id]
	if ok {
		return true
	}
	return false
}

func JoinRoom(room *models.Room, user models.User) error {
	doesRoomExist := Manager.DoesRoomExist(room)
	if !doesRoomExist {
		return fmt.Errorf("room doesnt exist")
	}
	room.Users = append(room.Users, user)

	return nil
}
