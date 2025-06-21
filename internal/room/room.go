package room

import (
	"fmt"
	"maps"
	"sync"
	"time"

	"github.com/JaanLavaerts/ticktickbrick/internal/data"
	"github.com/JaanLavaerts/ticktickbrick/internal/models"
	"github.com/JaanLavaerts/ticktickbrick/internal/util"
)

type RoomManager struct {
	rooms       map[string]*models.Room
	users_rooms map[string]string
	mu          sync.RWMutex
}

var Manager = &RoomManager{
	rooms:       make(map[string]*models.Room),
	users_rooms: make(map[string]string),
}

func CreateRoom(client *models.Client) (*models.Room, error) {
	team := data.RandomTeam()
	if otherRoom, err := Manager.GetRoomByUser(client.User); err == nil {
		return otherRoom, fmt.Errorf(util.UserAlreadyInRoomError)
	}

	clients := make(map[string]*models.Client)
	clients[client.User.Id] = client

	randomId := generateRoomId()
	room := &models.Room{
		Id:               randomId,
		Clients:          clients,
		CurrentTurn:      0, // TODO: should prob be random in the future, be ware of NextTurn logic/ tests
		TurnOrder:        []string{client.User.Id},
		CurrentTeam:      team,
		MentionedPlayers: nil,
		State:            models.RoomState(models.WAITING),
		StartTime:        time.Now(),
	}

	Manager.mu.Lock()
	defer Manager.mu.Unlock()

	client.Room = room
	Manager.AddRoom(room)
	Manager.users_rooms[client.User.Id] = room.Id
	return room, nil
}

func JoinRoom(room *models.Room, client *models.Client) error {
	room.Mu.Lock()
	defer room.Mu.Unlock()

	if room == nil || client.User.Id == "" {
		return fmt.Errorf(util.InvalidInputError)
	}

	if Manager.HasRoom(client.User) {
		return fmt.Errorf(util.UserAlreadyInRoomError)
	}

	client.Room = room
	room.TurnOrder = append(room.TurnOrder, client.User.Id)
	room.Clients[client.User.Id] = client
	Manager.mu.Lock()
	Manager.users_rooms[client.User.Id] = room.Id
	Manager.mu.Unlock()
	return nil
}

func SetRoomState(room *models.Room, state models.RoomState) {
	room.Mu.Lock()
	defer room.Mu.Unlock()

	room.State = state
}

func AllUsersReady(room *models.Room) bool {
	room.Mu.Lock()
	defer room.Mu.Unlock()

	allReady := true
	for _, client := range room.Clients {
		if !client.User.IsReady {
			allReady = false
			break
		}
	}
	return allReady
}

func (r *RoomManager) AddRoom(room *models.Room) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.rooms[room.Id] = room
}

func (r *RoomManager) GetRoom(id string) (*models.Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	room, ok := r.rooms[id]
	if !ok {
		return nil, fmt.Errorf(util.RoomNotFoundError)
	}
	return room, nil
}

func (r *RoomManager) GetAllRooms() (map[string]*models.Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	rooms := make(map[string]*models.Room)
	maps.Copy(rooms, r.rooms)
	if len(rooms) == 0 {
		return nil, fmt.Errorf(util.NoRoomsError)
	}
	return rooms, nil
}

func (r *RoomManager) GetRoomByUser(user models.User) (*models.Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

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
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, ok := r.users_rooms[user.Id]
	return ok
}

func (r *RoomManager) RoomExists(id string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, ok := r.rooms[id]
	return ok
}

func generateRoomId() string {
	return fmt.Sprintf("room_%d", time.Now().UnixNano())
}
