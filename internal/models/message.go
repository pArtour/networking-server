package models

type Message struct {
	Id           int64  `json:"id" validate:"required"`
	SenderId     int64  `json:"sender_id" validate:"required"`
	ReceiverId   int64  `json:"receiver_id" validate:"required"`
	ConnectionId int64  `json:"connection_id" validate:"required"`
	Content      string `json:"content" validate:"required"`
	//CreatedAt    string `json:"created_at" validate:"required"`
}

type CreateMessageInput struct {
	SenderId     int64  `json:"sender_id" validate:"required"`
	ReceiverId   int64  `json:"receiver_id" validate:"required"`
	ConnectionId int64  `json:"connection_id" validate:"required"`
	Content      string `json:"content" validate:"required"`
}
