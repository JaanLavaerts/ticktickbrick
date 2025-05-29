package game

type GameRoom struct {
	ID          string
	players     []User
	currentTeam string
	turnIndex   int
}
