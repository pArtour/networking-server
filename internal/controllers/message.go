package controllers

import (
	"github.com/pArtour/networking-server/internal/models"
	"github.com/pArtour/networking-server/internal/services"
)

type MessageController struct {
	service *services.MessageService
}

func NewMessageController(s *services.MessageService) *MessageController {
	return &MessageController{
		service: s,
	}
}

func (c *MessageController) CreateMessage(body *models.CreateMessageInput) (*models.Message, error) {
	return c.service.CreateMessage(body)
}

func (c *MessageController) GetMessagesForConnection(connectionId int64) ([]models.Message, error) {
	return c.service.GetMessagesByConnection(connectionId)
}
