package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pArtour/networking-server/internal/controllers"
	"github.com/pArtour/networking-server/internal/errors"
	"github.com/pArtour/networking-server/internal/helpers"
	"github.com/pArtour/networking-server/internal/models"
	"github.com/pArtour/networking-server/internal/validation"
)

type AuthHandler struct {
	controller *controllers.UserController
}

// LoginHandler is a function that handles the login route
func (h *AuthHandler) LoginHandler(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errors.ErrorResponse{Code: fiber.StatusBadRequest, Message: "Invalid input"})
	}

	// Validate input
	validationErrors := validation.ValidateStruct(input)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	user, err := h.controller.LoginUser(input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errors.ErrorResponse{Code: fiber.StatusBadRequest, Message: "Invalid credentials"})
	}

	// Create JWT token
	token, err := helpers.GenerateJWTToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: fmt.Sprintf("Error generating token: %s", err)})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

// RegisterHandler is a function that handles the register route
func (h *AuthHandler) RegisterHandler(c *fiber.Ctx) error {
	//type RegisterInput struct {
	//	Name           string `json:"name" validate:"required,min=2,max=100"`
	//	Email          string `json:"email" validate:"required,email"`
	//	Password       string `json:"password" validate:"required,min=8"`
	//	Bio            string `json:"bio" validate:"max=300"`
	//	ProfilePicture string `json:"profile_picture" validate:"url"`
	//}

	var input models.CreateUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errors.ErrorResponse{Code: fiber.StatusBadRequest, Message: "Invalid input"})
	}

	// Validate input
	validationErrors := validation.ValidateStruct(input)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	hashedPassword, err := helpers.HashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: fmt.Sprintf("Error hashing password: %s", err)})
	}

	input.Password = hashedPassword

	newUser, err := h.controller.CreateUser(&input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: fmt.Sprintf("Error creating user: %s", err)})
	}

	token, err := helpers.GenerateJWTToken(newUser.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: fmt.Sprintf("Error creating user: %s", err)})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func NewAuthHandler(router fiber.Router, uc *controllers.UserController) {
	h := &AuthHandler{controller: uc}
	authRouter := router.Group("/auth")
	authRouter.Post("/login", h.LoginHandler)
	authRouter.Post("/register", h.RegisterHandler)
}
