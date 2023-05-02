package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/pArtour/networking-server/internal/chat"
	"github.com/pArtour/networking-server/internal/controllers"
	"github.com/pArtour/networking-server/internal/errors"
	"github.com/pArtour/networking-server/internal/helpers"
	"github.com/pArtour/networking-server/internal/middleware"
	"github.com/pArtour/networking-server/internal/models"
	"strconv"
)

type ChatHandler struct {
	controller *controllers.MessageController
}

func NewChatHandler(router fiber.Router, c *controllers.MessageController) {
	messages := router.Group("/messages", middleware.JWTProtected())
	chat := router.Group("/chat", middleware.JWTProtected())
	h := &ChatHandler{
		controller: c,
	}
	chat.Post("/:connectionId", h.CreateMessageHandler, middleware.JWTProtected())
	chat.Get("/:connectionId", h.ChatHandler, middleware.JWTProtected(), websocket.New(h.WebSocketHandler))
	messages.Get("/:connectionId", h.GetChatHistoryHandler, middleware.JWTProtected())
}

func (h *ChatHandler) ChatHandler(c *fiber.Ctx) error {
	// Accept the WebSocket connection
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}

	return c.SendStatus(400)
}

func (h *ChatHandler) WebSocketHandler(c *websocket.Conn) {
	// Extract the connection ID from the URL
	connectionID, err := strconv.ParseInt(c.Params("connectionId"), 10, 64)
	if err != nil {
		c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "invalid connection ID"))
		c.Close()
		return
	}

	// Extract the user ID from the JWT token

	//userID, err := helpers.ExtractUserIDFromWebsocketJWT(c)
	//if err != nil {
	//	c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "invalid token"))
	//	c.Close()
	//	return
	//}

	// Add the WebSocket connection to the connections list
	chat.AddConnection(connectionID, c)
	defer chat.RemoveConnection(connectionID)

	// Read messages from the WebSocket and broadcast them to all other connected users
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			break
		}

		// Save the message in the database

		// Extract the receiver ID and content from the message
		var wsMessage models.ReceivedMessage
		err = json.Unmarshal(message, &wsMessage)
		if err != nil {
			// Handle JSON unmarshal error
			continue
		}
		receiverID := wsMessage.ReceiverId
		content := wsMessage.Message

		token := wsMessage.JWT
		userId, err := helpers.ExtractUserIDFromJWTString(token)
		if err != nil {
			c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "invalid token"))
			c.Close()
			break
		}
		_, err = h.controller.CreateMessage(&models.CreateMessageInput{
			SenderId:     userId,
			ReceiverId:   receiverID,
			ConnectionId: connectionID,
			Message:      content,
		})
		if err != nil {
			c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, fmt.Sprintf("failed to save message: %s", err.Error())))
			c.Close()
			break

		}

		chat.BroadcastMessage(connectionID, string(message))
	}
}

func (h *ChatHandler) GetChatHistoryHandler(c *fiber.Ctx) error {
	_, err := helpers.ExtractUserIDFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}
	connectionId, _ := strconv.ParseInt(c.Params("connectionId"), 10, 64)

	// Fetch the chat history from the database
	messages, err := h.controller.GetMessagesForConnection(connectionId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: fmt.Sprintf("Failed to fetch chat history: %s", err.Error())})
	}

	return c.JSON(messages)
}

func (h *ChatHandler) CreateMessageHandler(c *fiber.Ctx) error {
	// Extract the user ID from the JWT token
	userID, err := helpers.ExtractUserIDFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}

	// Extract the connection ID from the URL
	connectionID, err := strconv.ParseInt(c.Params("connectionId"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errors.ErrorResponse{Code: fiber.StatusBadRequest, Message: "Invalid connection ID"})
	}

	// Extract the receiver ID and content from the request body
	var body models.CreateMessageInput
	err = c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&errors.ErrorResponse{Code: fiber.StatusBadRequest, Message: "Invalid request body"})
	}

	// Save the message in the database
	message, err := h.controller.CreateMessage(&models.CreateMessageInput{
		SenderId:     userID,
		ReceiverId:   body.ReceiverId,
		ConnectionId: connectionID,
		Message:      body.Message,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: fmt.Sprintf("Failed to save message: %s", err.Error())})
	}

	return c.JSON(message)
}
