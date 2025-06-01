package room

import (
	"fmt"
	"maps"

	"github.com/JaanLavaerts/ticktickbrick/internal/game"
)

type RoomManager struct {
	rooms map[string]*game.Room
}

var Manager = &RoomManager{
	rooms: make(map[string]*game.Room),
}

func (r *RoomManager) AddRoom(room *game.Room) {
	r.rooms[room.Id] = room
}

func (r *RoomManager) GetRoom(id string) *game.Room {
	return r.rooms[id]
}

func (r *RoomManager) GetAllRooms() map[string]*game.Room {
	rooms := make(map[string]*game.Room)
	maps.Copy(rooms, r.rooms)
	return rooms
}

func (r *RoomManager) DoesRoomExist(room *game.Room) bool {
	_, ok := r.rooms[room.Id]
	if ok {
		return true
	}
	return false
}

func JoinRoom(room *game.Room, user game.User) error {
	doesRoomExist := Manager.DoesRoomExist(room)
	if !doesRoomExist {
		return fmt.Errorf("room doesnt exist")
	}
	room.Users = append(room.Users, user)

	return nil
}
