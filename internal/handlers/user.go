package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pArtour/networking-server/internal/controllers"
	"strconv"
)

type UserHandler struct {
	controller *controllers.UserController
}

func NewUserHandler(app *fiber.App, uc *controllers.UserController) {
	usersRouter := app.Group("/users")
	uh := &UserHandler{
		controller: uc,
	}

	uh.setupUserRoutes(usersRouter)
}

func (uh *UserHandler) setupUserRoutes(r fiber.Router) {
	r.Get("/users", uh.getUsersHandler)
	r.Get("/users/:id", uh.getUserHandler)
	r.Post("/users", uh.createUserHandler)
	r.Put("/users/:id", uh.updateUserHandler)
	r.Delete("/users/:id", uh.deleteUserHandler)
}

func (uh *UserHandler) getUsersHandler(c *fiber.Ctx) error {
	users, err := uh.controller.GetUsers()
	if err != nil {
		return c.Status(500).SendString("Error fetching users")
	}
	return c.JSON(users)
}

func (uh *UserHandler) getUserHandler(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid handlers ID")
	}
	user, err := uh.controller.GetUserById(id)
	if err != nil {
		return c.Status(500).SendString("Error fetching handlers")
	}
	return c.JSON(user)

}

func (uh *UserHandler) createUserHandler(c *fiber.Ctx) error {
	name := c.FormValue("name")
	if name == "" {
		return c.Status(400).SendString("Name is required")
	}
	id, err := uh.controller.CreateUser(name)
	if err != nil {
		return c.Status(500).SendString("Error creating handlers")
	}
	return c.JSON(fiber.Map{"id": id})
}

func (uh *UserHandler) updateUserHandler(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid handlers ID")
	}
	name := c.FormValue("name")
	if name == "" {
		return c.Status(400).SendString("Name is required")
	}
	err = uh.controller.UpdateUser(id, name)
	if err != nil {
		return c.Status(500).SendString("Error updating handlers")
	}
	return c.SendStatus(200)
}

func (uh *UserHandler) deleteUserHandler(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid handlers ID")
	}
	err = uh.controller.DeleteUser(id)
	if err != nil {
		return c.Status(500).SendString("Error deleting handlers")
	}
	return c.SendStatus(200)
}
