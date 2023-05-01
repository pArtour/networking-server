package services

import (
	"context"
	"github.com/pArtour/networking-server/internal/database"
	"github.com/pArtour/networking-server/internal/models"
)

type MessageService struct {
	db *database.Db
}

func NewMessageService(db *database.Db) *MessageService {
	return &MessageService{
		db: db,
	}
}

// GetMessagesByConnection returns all messages between two users
func (s *MessageService) GetMessagesByConnection(connectionId int64) ([]models.Message, error) {
	rows, err := s.db.Conn.Query(context.Background(), "SELECT id, sender_id, receiver_id, connection_id, message FROM messages WHERE connection_id = $1", connectionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message
		err := rows.Scan(&message.Id, &message.SenderId, &message.ReceiverId, &message.Message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

// CreateMessage creates a new message and returns it
func (s *MessageService) CreateMessage(body *models.CreateMessageInput) (*models.Message, error) {
	var message models.Message
	err := s.db.Conn.QueryRow(context.Background(), "INSERT INTO messages (sender_id, receiver_id, message) VALUES ($1, $2, $3) RETURNING id, sender_id, receiver_id, message", body.SenderId, body.ReceiverId, body.Message).Scan(&message.Id, &message.SenderId, &message.ReceiverId, &message.Message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
