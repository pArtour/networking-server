package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pArtour/networking-server/internal/controllers"
	"github.com/pArtour/networking-server/internal/errors"
	"github.com/pArtour/networking-server/internal/models"
	"github.com/pArtour/networking-server/internal/validation"
	"strconv"
)

// UserHandler is a struct that contains all handlers for users
type UserHandler struct {
	controller *controllers.UserController
}

// NewUserHandler returns a new UserHandler struct
func NewUserHandler(router fiber.Router, uc *controllers.UserController) {
	usersRouter := router.Group("/users")
	uh := &UserHandler{
		controller: uc,
	}

	uh.setupUserRoutes(usersRouter)
}

// setupUserRoutes sets up all routes for users
func (uh *UserHandler) setupUserRoutes(r fiber.Router) {
	r.Get("/", uh.getUsersHandler)
	r.Get("/:id", uh.getUserHandler)
	r.Post("/", uh.createUserHandler)
	r.Put("/:id", uh.updateUserHandler)
	r.Delete("/:id", uh.deleteUserHandler)
}

// getUsersHandler handles GET /users
func (uh *UserHandler) getUsersHandler(c *fiber.Ctx) error {
	users, err := uh.controller.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: "Error fetching users"})
	}
	return c.JSON(users)
}

// getUserHandler handles GET /users/:id
func (uh *UserHandler) getUserHandler(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errors.ErrorResponse{Code: fiber.StatusBadRequest, Message: "Invalid user ID"})
	}
	user, err := uh.controller.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: "Error fetching user"})
	}
	return c.JSON(user)

}

// createUserHandler handles POST /users
func (uh *UserHandler) createUserHandler(c *fiber.Ctx) error {
	user := new(models.CreateUserBody)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errors.ErrorResponse{Code: fiber.StatusBadRequest, Message: "Invalid request body"})
	}

	validationErrors := validation.ValidateStruct(*user)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	id, err := uh.controller.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: "Error creating user"})
	}
	return c.JSON(fiber.Map{"id": id})
}

// updateUserHandler handles PUT /users/:id
func (uh *UserHandler) updateUserHandler(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errors.ErrorResponse{Code: fiber.StatusBadRequest, Message: "Invalid user ID"})
	}
	user := new(models.UpdateUserBody)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errors.ErrorResponse{Code: fiber.StatusBadRequest, Message: "Invalid request body"})
	}

	validationErrors := validation.ValidateStruct(*user)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	err = uh.controller.UpdateUser(id, *user)
	if err != nil {
		return c.Status(500).SendString("Error updating handlers")
	}
	return c.SendStatus(fiber.StatusOK)
}

// deleteUserHandler handles DELETE /users/:id
func (uh *UserHandler) deleteUserHandler(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid handlers ID")
	}
	err = uh.controller.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting handlers")
	}
	return c.SendStatus(fiber.StatusOK)
}
