package controllers

import (
	"github.com/pArtour/networking-server/internal/models"
	"github.com/pArtour/networking-server/internal/services"
)

// UserController is a struct that contains a UserService
type UserController struct {
	userService *services.UserService
	authService *services.AuthService
}

// NewUserController returns a new UserController struct
func NewUserController(us *services.UserService, as *services.AuthService) *UserController {
	return &UserController{
		userService: us,
		authService: as,
	}
}

// GetUsers returns all users
func (c *UserController) GetUsers() ([]models.User, error) {
	users, err := c.userService.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserById returns a user by id
func (c *UserController) GetUserById(id int64) (*models.User, error) {
	user, err := c.userService.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *UserController) LoginUser(email, password string) (*models.User, error) {
	user, err := c.userService.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	err = c.authService.CheckUserCredentials(password, user)
	if err != nil {
		return nil, err
	}
	return &models.User{Name: user.Name, ID: user.ID, Email: user.Email}, nil
}

// CreateUser creates a new user
func (c *UserController) CreateUser(body *models.CreateUserInput) (*models.User, error) {
	user, err := c.userService.CreateUser(body)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates a user
func (c *UserController) UpdateUser(id int64, body *models.UpdateUserInput) error {
	err := c.userService.UpdateUser(id, body)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user
func (c *UserController) DeleteUser(id int64) error {
	err := c.userService.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
