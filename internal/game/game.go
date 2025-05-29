package game

type GameRoom struct {
	Id          string
	Players     []User
	CurrentTeam string
	TurnIndex   int
}
