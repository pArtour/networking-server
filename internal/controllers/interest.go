package controllers

import (
	"github.com/pArtour/networking-server/internal/models"
	"github.com/pArtour/networking-server/internal/services"
)

type InterestController struct {
	interestService *services.InterestService
}

func NewInterestController(is *services.InterestService) *InterestController {
	return &InterestController{
		interestService: is,
	}
}

// GetUserInterests returns all interests for a user
func (c *InterestController) GetUserInterests(id int64) ([]models.Interest, error) {
	interests, err := c.interestService.GetInterestsByUserId(id)
	if err != nil {
		return nil, err
	}
	return interests, nil
}

// CreateInterest creates a new interest
func (c *InterestController) CreateInterest(body *models.CreateInterestInput) (*models.Interest, error) {
	interest, err := c.interestService.CreateInterest(body)
	if err != nil {
		return nil, err
	}
	return interest, nil
}

// GetInterests returns all interests
func (c *InterestController) GetInterests() ([]models.Interest, error) {
	interests, err := c.interestService.GetInterests()
	if err != nil {
		return nil, err
	}
	return interests, nil
}

// AddUserInterest adds an interest to a user
func (c *InterestController) AddUserInterest(userId int64, interestId int64) error {
	err := c.interestService.AddInterestToUser(&models.AddInterestToUserInput{
		UserId:     userId,
		InterestId: interestId,
	})

	if err != nil {
		return err
	}
	return nil
}

func (c *InterestController) DeleteUserInterest(body *models.DeleteInterestFromUserInput) error {
	err := c.interestService.DeleteInterestFromUser(body.UserId, body.InterestId)

	if err != nil {
		return err
	}
	return nil
}
