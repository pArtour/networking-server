package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pArtour/networking-server/internal/controllers"
	"github.com/pArtour/networking-server/internal/errors"
	"github.com/pArtour/networking-server/internal/helpers"
	"github.com/pArtour/networking-server/internal/middleware"
	"github.com/pArtour/networking-server/internal/models"
)

type ConnectionHandler struct {
	controller *controllers.ConnectionController
}

func NewConnectionHandler(router fiber.Router, controller *controllers.ConnectionController) {
	handler := &ConnectionHandler{
		controller: controller,
	}
	connectionsRouter := router.Group("/connections", middleware.JWTProtected())
	connectionsRouter.Get("/", handler.GetUserConnections)
	connectionsRouter.Post("/", handler.CreateConnection)
	connectionsRouter.Delete("/:id", handler.DeleteConnection)
	connectionsRouter.Get("/:id", handler.GetConnectionById)
}

// GetUserConnections returns all connections for a user
func (h *ConnectionHandler) GetUserConnections(c *fiber.Ctx) error {
	userId, err := helpers.ExtractUserIDFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}
	connections, err := h.controller.GetUserConnections(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(connections)
}

// CreateConnection creates a new connection
func (h *ConnectionHandler) CreateConnection(c *fiber.Ctx) error {
	userId, err := helpers.ExtractUserIDFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}

	body := new(models.CreateConnectionInput)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}
	connection, err := h.controller.CreateConnection(&models.CreateConnectionRecordInput{
		UserId:       userId,
		TargetUserId: body.ConnectWithUser,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(connection)
}

// DeleteConnection deletes a connection
func (h *ConnectionHandler) DeleteConnection(c *fiber.Ctx) error {
	userId, err := helpers.ExtractUserIDFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errors.ErrorResponse{Code: fiber.StatusBadRequest, Message: err.Error()})
	}
	err = h.controller.DeleteUserConnection(int64(id), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}

// GetConnectionById returns a connection by id
func (h *ConnectionHandler) GetConnectionById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errors.ErrorResponse{Code: fiber.StatusBadRequest, Message: err.Error()})
	}
	connection, err := h.controller.GetConnectionById(int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(connection)
}
