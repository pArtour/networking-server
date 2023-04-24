package services

import "github.com/pArtour/networking-server/internal/database"

// Services is a struct that contains all services
type Services struct {
	Us *UserService
	As *AuthService
	Is *InterestService
	Cs *ConnectionService
	Ms *MessageService
}

// NewServices returns a new Services struct
func NewServices(Db *database.Db) *Services {
	return &Services{
		Us: NewUserService(Db),
		As: NewAuthService(Db),
		Is: NewInterestService(Db),
		Cs: NewConnectionService(Db),
		Ms: NewMessageService(Db),
	}
}
