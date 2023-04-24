package models

type Connection struct {
	Id        int64  `json:"id" validate:"required"`
	UserId1   int64  `json:"user_id_1" validate:"required"`
	UserId2   int64  `json:"user_id_2" validate:"required"`
	CreatedAt string `json:"created_at" validate:"required"`
}

type CreateConnectionInput struct {
	ConnectWithUser int64 `json:"connect_with_user" validate:"required"`
}

type CreateConnectionRecordInput struct {
	UserId       int64 `json:"user_id_1" validate:"required"`
	TargetUserId int64 `json:"user_id_2" validate:"required"`
}
