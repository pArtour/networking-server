package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pArtour/networking-server/internal/controllers"
	"github.com/pArtour/networking-server/internal/errors"
	"github.com/pArtour/networking-server/internal/helpers"
	"github.com/pArtour/networking-server/internal/middleware"
	"github.com/pArtour/networking-server/internal/models"
)

type InterestHandler struct {
	interestController *controllers.InterestController
}

func NewInterestHandler(router fiber.Router, ic *controllers.InterestController) {
	handler := &InterestHandler{
		interestController: ic,
	}
	interestRouter := router.Group("/interests", middleware.JWTProtected())
	interestRouter.Get("/", handler.GetInterests)
	interestRouter.Get("/me", handler.GetUserInterests)
	interestRouter.Post("/:id", handler.AddInterest)
	interestRouter.Delete("/:id", handler.DeleteInterest)
}

// GetCurrentUserInterests returns all interests for a user
func (h *InterestHandler) GetUserInterests(c *fiber.Ctx) error {
	userId, err := helpers.ExtractUserIDFromJWT(c)
	interests, err := h.interestController.GetUserInterests(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(interests)
}

// GetInterests returns all interests
func (h *InterestHandler) GetInterests(c *fiber.Ctx) error {
	interests, err := h.interestController.GetInterests()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(interests)
}

// AddInterest adds an interest to a user
func (h *InterestHandler) AddInterest(c *fiber.Ctx) error {
	userId, err := helpers.ExtractUserIDFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errors.ErrorResponse{Code: fiber.StatusBadRequest, Message: err.Error()})
	}
	interestId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errors.ErrorResponse{Code: fiber.StatusBadRequest, Message: err.Error()})
	}
	err = h.interestController.AddUserInterest(userId, int64(interestId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}

// DeleteInterest deletes an interest from a user
func (h *InterestHandler) DeleteInterest(c *fiber.Ctx) error {
	userId, err := helpers.ExtractUserIDFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errors.ErrorResponse{Code: fiber.StatusBadRequest, Message: err.Error()})
	}
	interestId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errors.ErrorResponse{Code: fiber.StatusBadRequest, Message: err.Error()})
	}
	err = h.interestController.DeleteUserInterest(&models.DeleteInterestFromUserInput{UserId: userId, InterestId: int64(interestId)})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}
