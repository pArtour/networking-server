package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pArtour/networking-server/internal/controllers"
)

type Handlers struct {
	Uh *UserHandler
}

func NewHandlers(app *fiber.App, c *controllers.Controllers) {
	NewUserHandler(app, c.UserController)
}
