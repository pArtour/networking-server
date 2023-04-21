package models

// User is a struct that contains a user's id and name
type User struct {
	ID   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,min=1,max=32"`
}

type CreateUserBody struct {
	Name string `json:"name" validate:"required,min=1,max=32"`
}

type UpdateUserBody struct {
	Name string `json:"name" validate:"min=1,max=32"`
}
