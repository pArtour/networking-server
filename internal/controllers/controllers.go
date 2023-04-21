package controllers

import "github.com/pArtour/networking-server/internal/services"

type Controllers struct {
	UserController *UserController
}

func NewControllers(s *services.Services) *Controllers {
	return &Controllers{
		UserController: NewUserController(s.Us),
	}
}
