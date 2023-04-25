package services

import (
	"context"
	"github.com/pArtour/networking-server/internal/database"
	"github.com/pArtour/networking-server/internal/models"
)

type InterestService struct {
	db *database.Db
}

func NewInterestService(db *database.Db) *InterestService {
	return &InterestService{
		db: db,
	}
}

// GetInterests returns all interests
func (s *InterestService) GetInterests() ([]models.Interest, error) {
	rows, err := s.db.Conn.Query(context.Background(), "SELECT id, name FROM interests")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var interests []models.Interest
	for rows.Next() {
		var interest models.Interest
		err := rows.Scan(&interest.Id, &interest.Name)
		if err != nil {
			return nil, err
		}
		interests = append(interests, interest)
	}

	return interests, nil
}

// CreateInterest creates a new interest and returns it
func (s *InterestService) CreateInterest(body *models.CreateInterestInput) (*models.Interest, error) {
	var interest models.Interest
	err := s.db.Conn.QueryRow(context.Background(), "INSERT INTO interests (name) VALUES ($1) RETURNING id, name", body.Name).Scan(&interest.Id, &interest.Name)
	if err != nil {
		return nil, err
	}
	return &interest, nil
}

// GetInterestsByUserId returns all interests of a user
func (s *InterestService) GetInterestsByUserId(userId int64) ([]models.Interest, error) {
	rows, err := s.db.Conn.Query(context.Background(), "SELECT i.id, i.name FROM interests i JOIN user_interests ui ON i.id = ui.interest_id WHERE ui.user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var interests []models.Interest
	for rows.Next() {
		var interest models.Interest
		err := rows.Scan(&interest.Id, &interest.Name)
		if err != nil {
			return nil, err
		}
		interests = append(interests, interest)
	}

	return interests, nil
}

// AddInterestToUser adds an interest to a user
func (s *InterestService) AddInterestToUser(body *models.AddInterestToUserInput) error {
	_, err := s.db.Conn.Exec(context.Background(), "INSERT INTO user_interests (user_id, interest_id) VALUES ($1, $2) RETURNING id", body.UserId, body.InterestId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteInterestFromUser deletes an interest from a user
func (s *InterestService) DeleteInterestFromUser(userId int64, interestId int64) error {
	_, err := s.db.Conn.Exec(context.Background(), "DELETE FROM user_interests WHERE user_id = $1 AND interest_id = $2", userId, interestId)
	if err != nil {
		return err
	}
	return nil
}
