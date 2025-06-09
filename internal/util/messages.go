package util

const (
	// error
	UserAlreadyInRoomError  = "user is already in a room"
	UserNotInRoomError      = "user is not in the room"
	UserDoesntHaveRoomError = "user doesnt have a room"
	RoomNotFoundError       = "room doesn't exist"
	NoRoomsError            = "no rooms exist"
	InvalidInputError       = "invalid input"
	GameNotOverYetError     = "game is not over yet"

	// success
	UserJoinedRoomSuccess = "user joined room succesfully"
	RoomCreatedSuccess    = "room created succesfully"
	RoomsRetrievedSuccess = "rooms retrieved succesfully"
)
