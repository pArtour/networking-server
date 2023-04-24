package models

type Interest struct {
	Id   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,min=1,max=32"`
}

type UpdateInterestInput struct {
	Name string `json:"name" validate:"required,min=1,max=32"`
}

type CreateInterestInput struct {
	Name string `json:"name" validate:"required,min=1,max=32"`
}

type AddInterestToUserInput struct {
	UserId     int64 `json:"user_id" validate:"required"`
	InterestId int64 `json:"interest_id" validate:"required"`
}

type DeleteInterestFromUserInput struct {
	UserId     int64 `json:"user_id" validate:"required"`
	InterestId int64 `json:"interest_id" validate:"required"`
}
