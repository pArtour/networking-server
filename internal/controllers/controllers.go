package controllers

import "github.com/pArtour/networking-server/internal/services"

// Controllers is a struct that contains all controllers
type Controllers struct {
	UserController       *UserController
	ConnectionController *ConnectionController
	InterestController   *InterestController
	MessageController    *MessageController
}

// NewControllers returns a new Controllers struct
func NewControllers(s *services.Services) *Controllers {
	return &Controllers{
		UserController:       NewUserController(s.Us, s.As),
		ConnectionController: NewConnectionController(s.Cs),
		InterestController:   NewInterestController(s.Is),
		MessageController:    NewMessageController(s.Ms),
	}
}
