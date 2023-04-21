package controllers

import (
	"github.com/pArtour/networking-server/internal/models"
	"github.com/pArtour/networking-server/internal/services"
)

// UserController is a struct that contains a UserService
type UserController struct {
	service *services.UserService
}

// NewUserController returns a new UserController struct
func NewUserController(us *services.UserService) *UserController {
	return &UserController{
		service: us,
	}
}

// GetUsers returns all users
func (uc *UserController) GetUsers() ([]models.User, error) {
	users, err := uc.service.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserById returns a user by id
func (uc *UserController) GetUserById(id int64) (models.User, error) {
	user, err := uc.service.GetUser(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// CreateUser creates a new user
func (uc *UserController) CreateUser(name string) (int64, error) {
	id, err := uc.service.CreateUser(name)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateUser updates a user
func (uc *UserController) UpdateUser(id int64, name string) error {
	err := uc.service.UpdateUser(id, name)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user
func (uc *UserController) DeleteUser(id int64) error {
	err := uc.service.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
