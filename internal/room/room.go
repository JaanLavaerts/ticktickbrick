package room

import "github.com/JaanLavaerts/ticktickbrick/internal/game"

type RoomManager struct {
	rooms map[string]*game.Room
}

var roomManager = &RoomManager{
	rooms: make(map[string]*game.Room),
}

func (r *RoomManager) AddRoom(room *game.Room) {
	roomManager.rooms[room.Id] = room
}

func (r *RoomManager) GetRoom(id string) *game.Room {
	return roomManager.rooms[id]
}

func (r *RoomManager) GetAllRooms() map[string]*game.Room {
	rooms := make(map[string]*game.Room)
	for key, val := range roomManager.rooms {
		rooms[key] = val
	}
	return rooms
}
