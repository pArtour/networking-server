package services

import "github.com/pArtour/networking-server/internal/database"

type Services struct {
	Us *UserService
}

func NewServices(Db *database.Db) *Services {
	return &Services{
		Us: NewUserService(Db),
	}
}
