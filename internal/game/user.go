package game

type User struct {
	Id          string
	Username    string
	Lives       int
	HasAnswered bool // has the player answered yet this round
}
