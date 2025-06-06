package room

import (
	"fmt"
	"maps"
	"time"

	"github.com/JaanLavaerts/ticktickbrick/internal/models"
	"github.com/JaanLavaerts/ticktickbrick/internal/util"
)

type RoomManager struct {
	rooms       map[string]*models.Room
	users_rooms map[string]string
}

var Manager = &RoomManager{
	rooms:       make(map[string]*models.Room),
	users_rooms: make(map[string]string), // map of user_id:room_id to keep track of which user is in what room
}

func CreateRoom(user models.User, team models.Team) (models.Room, error) {
	if otherRoom, err := Manager.GetRoomByUser(user); err == nil {
		return *otherRoom, fmt.Errorf(util.UserInRoomError)
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
	Manager.rooms[room.Id] = room
	Manager.users_rooms[user.Id] = room.Id
	return *room, nil
}

func JoinRoom(room *models.Room, user models.User) error {
	if room == nil || user.Id == "" {
		return fmt.Errorf(util.InvalidInputError)
	}

	if Manager.HasRoom(user) {
		return fmt.Errorf(util.UserInRoomError)
	}

	if !Manager.RoomExists(room.Id) {
		return fmt.Errorf(util.RoomNotFoundError)
	}

	room.Users = append(room.Users, user)
	Manager.users_rooms[user.Id] = room.Id

	return nil
}

func (r *RoomManager) AddRoom(room *models.Room) {
	r.rooms[room.Id] = room
}

func (r *RoomManager) GetRoom(id string) (*models.Room, error) {
	room, ok := r.rooms[id]
	if !ok {
		return nil, fmt.Errorf(util.RoomNotFoundError)
	}
	return room, nil
}

func (r *RoomManager) GetAllRooms() (map[string]*models.Room, error) {
	rooms := make(map[string]*models.Room)
	maps.Copy(rooms, r.rooms)
	if len(rooms) == 0 {
		return nil, fmt.Errorf(util.NoRoomsError)
	}
	return rooms, nil
}

func (r *RoomManager) GetRoomByUser(user models.User) (*models.Room, error) {
	id, ok := r.users_rooms[user.Id]
	if !ok {
		return nil, fmt.Errorf(util.UserDoesntHaveRoomError)
	}
	room, err := r.GetRoom(id)
	if err != nil {
		return nil, fmt.Errorf(util.RoomNotFoundError)
	}
	return room, nil
}

func (r *RoomManager) HasRoom(user models.User) bool {
	_, ok := r.users_rooms[user.Id]
	return ok
}

func (r *RoomManager) RoomExists(id string) bool {
	_, ok := r.rooms[id]
	return ok
}

func generateTimestampID() string {
	return fmt.Sprintf("room_%d", time.Now().UnixNano())
}
