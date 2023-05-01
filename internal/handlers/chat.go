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

	chat.Get("/", h.ChatHandler, middleware.JWTProtected(), websocket.New(h.WebSocketHandler))
	messages.Get("/:userId", h.GetChatHistoryHandler, middleware.JWTProtected())
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
	// Extract the user ID from the JWT token
	userID, err := helpers.ExtractUserIDFromWebsocketJWT(c)
	if err != nil {
		c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "invalid token"))
		c.Close()
		return
	}

	// Add the WebSocket connection to the connections list
	chat.AddConnection(userID, c)
	defer chat.RemoveConnection(userID)

	// Read messages from the WebSocket and broadcast them to all other connected users
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			break
		}

		// Save the message in the database
		senderID := userID

		// Extract the receiver ID and content from the message
		var wsMessage models.CreateMessageInput
		err = json.Unmarshal(message, &wsMessage)
		if err != nil {
			// Handle JSON unmarshal error
			continue
		}
		receiverID := wsMessage.ReceiverId
		content := wsMessage.Message

		_, err = h.controller.CreateMessage(&models.CreateMessageInput{
			SenderId:   senderID,
			ReceiverId: receiverID,
			Message:    content,
		})
		if err != nil {
			// Handle database error
			continue
		}

		chat.BroadcastMessage(userID, string(message))
	}
}

func (h *ChatHandler) GetChatHistoryHandler(c *fiber.Ctx) error {
	userID1, err := helpers.ExtractUserIDFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: err.Error()})
	}
	userID2, _ := strconv.ParseInt(c.Params("userId"), 10, 64)

	// Fetch the chat history from the database
	messages, err := h.controller.GetMessagesBetweenUsers(userID1, userID2)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.ErrorResponse{Code: fiber.StatusInternalServerError, Message: fmt.Sprintf("Failed to fetch chat history: %s", err.Error())})
	}

	return c.JSON(messages)
}
