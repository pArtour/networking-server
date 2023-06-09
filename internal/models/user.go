package models

// User is a struct that contains a user's id and name
type User struct {
	ID             int64  `json:"id" validate:"required"`
	Name           string `json:"name" validate:"required,min=1,max=32"`
	Email          string `json:"email" validate:"required,email"`
	Bio            string `json:"bio" validate:"required,min=25"`
	ProfilePicture string `json:"profile_picture"`
}

type UserWithPassword struct {
	User
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterUserInputInterest struct {
	ID int64 `json:"id" validate:"required"`
}

type CreateUserInput struct {
	Name           string `json:"name" validate:"required,min=1,max=32"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,min=8"`
	ProfilePicture string `json:"profile_picture"`
	Bio            string `json:"bio" validate:"required,min=25"`
}

type UpdateUserInput struct {
	Name           string `json:"name" validate:"min=1,max=32"`
	Email          string `json:"email" validate:"email"`
	Bio            string `json:"bio" validate:"required,min=25"`
	ProfilePicture string `json:"profile_picture"`
}

type UserWithInterests struct {
	User
	Interests []Interest `json:"interests"`
}
