package room

import (
	"fmt"
	"maps"
	"time"

	"github.com/JaanLavaerts/ticktickbrick/internal/models"
)

type RoomManager struct {
	rooms       map[string]*models.Room
	users_rooms map[string]string
}

var Manager = &RoomManager{
	rooms:       make(map[string]*models.Room),
	users_rooms: make(map[string]string),
}

func CreateRoom(user models.User, team models.Team) (models.Room, error) {
	otherRoom := Manager.GetRoomByUser(user)
	if otherRoom != nil {
		return *otherRoom, fmt.Errorf("user already has a room")
	}
	randomId := generateTimestampID()
	room := &models.Room{
		Id:               randomId,
		Users:            []models.User{user},
		CurrentTurn:      0, // TODO: should be random in the future, be ware of NextTurn logic/ tests
		CurrentTeam:      team,
		MentionedPlayers: nil,
		State:            models.RoomState(models.INPROGRESS),
		StartTime:        time.Now(),
	}
	Manager.users_rooms[user.Id] = room.Id
	return *room, nil
}

func JoinRoom(room *models.Room, user models.User) error {
	doesRoomExist := Manager.DoesRoomExist(room)
	if !doesRoomExist {
		return fmt.Errorf("room doesnt exist")
	}
	room.Users = append(room.Users, user)

	return nil
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

func (r *RoomManager) GetRoomByUser(user models.User) *models.Room {
	id, ok := r.users_rooms[user.Id]
	if !ok {
		return nil
	}
	return r.GetRoom(id)
}

func (r *RoomManager) DoesRoomExist(room *models.Room) bool {
	_, ok := r.rooms[room.Id]
	if ok {
		return true
	}
	return false
}

func generateTimestampID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
