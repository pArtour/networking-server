package services

import (
	"context"
	"github.com/pArtour/networking-server/internal/database"
	"github.com/pArtour/networking-server/internal/models"
)

type UserService struct {
	db *database.Db
}

func NewUserService(db *database.Db) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	rows, err := s.db.Conn.Query(context.Background(), "SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *UserService) CreateUser(name string) (int64, error) {
	var id int64
	err := s.db.Conn.QueryRow(context.Background(), "INSERT INTO users (name) VALUES ($1) RETURNING id", name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *UserService) UpdateUser(id int64, name string) error {
	_, err := s.db.Conn.Exec(context.Background(), "UPDATE users SET name=$1 WHERE id=$2", name, id)
	return err
}

func (s *UserService) DeleteUser(id int64) error {
	_, err := s.db.Conn.Exec(context.Background(), "DELETE FROM users WHERE id=$1", id)
	return err
}

func (s *UserService) GetUser(id int64) (models.User, error) {
	var user models.User
	err := s.db.Conn.QueryRow(context.Background(), "SELECT id, name FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
