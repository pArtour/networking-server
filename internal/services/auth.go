package services

import (
	"errors"
	"github.com/pArtour/networking-server/internal/database"
	"github.com/pArtour/networking-server/internal/helpers"
	"github.com/pArtour/networking-server/internal/models"
)

type AuthService struct {
	db *database.Db
}

func NewAuthService(db *database.Db) *AuthService {
	return &AuthService{db: db}
}

func (s AuthService) CheckUserCredentials(password string, user *models.UserWithPassword) error {
	if !helpers.CheckPasswordHash(password, user.Password) {
		return errors.New("invalid credentials")
	}
	return nil
}
