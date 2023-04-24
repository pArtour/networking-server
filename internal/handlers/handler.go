package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pArtour/networking-server/internal/controllers"
)

// Handlers is a struct that contains all handlers
type Handlers struct {
	Uh *UserHandler
	Ah *AuthHandler
	Ch *ConnectionHandler
	Ic *InterestHandler
}

// NewHandlers returns a new Handlers struct
func NewHandlers(router fiber.Router, c *controllers.Controllers) {
	NewUserHandler(router, c.UserController)
	NewAuthHandler(router, c.UserController)
	NewConnectionHandler(router, c.ConnectionController)
	NewInterestHandler(router, c.InterestController)
}
