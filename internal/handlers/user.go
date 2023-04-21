package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pArtour/networking-server/internal/controllers"
	"github.com/pArtour/networking-server/internal/errors"
	"github.com/pArtour/networking-server/internal/helpers"
	"github.com/pArtour/networking-server/internal/models"
	"github.com/pArtour/networking-server/internal/validation"
)

// UserHandler is a struct that contains all handlers for users
type UserHandler struct {
	controller *controllers.UserController
}

// NewUserHandler returns a new UserHandler struct
func NewUserHandler(router fiber.Router, uc *controllers.UserController) {
	usersRouter := router.Group("/users")
	h := &UserHandler{
		controller: uc,
	}

	h.setupUserRoutes(usersRouter)
}

// setupUserRoutes sets up all routes for users
func (h *UserHandler) setupUserRoutes(r fiber.Router) {
	r.Get("/", h.getUsersHandler)
	r.Get("/me", h.getCurrentUserHandler)
	r.Put("/", h.updateUserHandler)
	r.Delete("/", h.deleteUserHandler)
}

// getUsersHandler handles GET /users
func (h *UserHandler) getUsersHandler(c *fiber.Ctx) error {
	users, err := h.controller.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: "Error fetching users"})
	}
	return c.JSON(users)
}

// registerUserHandler handles POST /users
func (h *UserHandler) registerUserHandler(c *fiber.Ctx) error {
	user := new(models.CreateUserInput)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errors.ErrorResponse{Code: fiber.StatusBadRequest, Message: "Invalid request body"})
	}

	validationErrors := validation.ValidateStruct(*user)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	newUser, err := h.controller.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: "Error creating user"})
	}
	return c.JSON(newUser)
}

// updateUserHandler handles PUT /users/:id
func (h *UserHandler) updateUserHandler(c *fiber.Ctx) error {
	id, err := helpers.ExtractUserIDFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&errors.ErrorResponse{Code: fiber.StatusUnauthorized, Message: "Unauthorized"})
	}
	user := new(models.UpdateUserInput)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errors.ErrorResponse{Code: fiber.StatusBadRequest, Message: "Invalid request body"})
	}

	validationErrors := validation.ValidateStruct(*user)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	_, err = h.controller.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&errors.ErrorResponse{Code: fiber.StatusNotFound, Message: "User not found"})
	}

	err = h.controller.UpdateUser(id, user)
	if err != nil {
		return c.Status(500).SendString("Error updating handlers")
	}
	return c.SendStatus(fiber.StatusOK)
}

// deleteUserHandler handles DELETE /users/:id
func (h *UserHandler) deleteUserHandler(c *fiber.Ctx) error {
	id, err := helpers.ExtractUserIDFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&errors.ErrorResponse{Code: fiber.StatusUnauthorized, Message: "Unauthorized"})
	}
	err = h.controller.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting handlers")
	}
	return c.SendStatus(fiber.StatusOK)
}

// getCurrentUserHandler handles GET /users/:id (for the current user) and returns the current user
func (h *UserHandler) getCurrentUserHandler(c *fiber.Ctx) error {
	userId, err := helpers.ExtractUserIDFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&errors.ErrorResponse{Code: fiber.StatusUnauthorized, Message: "Unauthorized"})
	}

	user, err := h.controller.GetUserById(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: "Error fetching user"})
	}

	return c.JSON(user)
}
