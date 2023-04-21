package controllers

import (
	"github.com/pArtour/networking-server/internal/models"
	"github.com/pArtour/networking-server/internal/services"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(us *services.UserService) *UserController {
	return &UserController{
		service: us,
	}
}

func (uc *UserController) GetUsers() ([]models.User, error) {
	users, err := uc.service.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *UserController) GetUserById(id int64) (models.User, error) {
	user, err := uc.service.GetUser(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (uc *UserController) CreateUser(name string) (int64, error) {
	id, err := uc.service.CreateUser(name)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (uc *UserController) UpdateUser(id int64, name string) error {
	err := uc.service.UpdateUser(id, name)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserController) DeleteUser(id int64) error {
	err := uc.service.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
