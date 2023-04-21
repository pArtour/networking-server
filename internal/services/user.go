package services

import (
	"context"
	"github.com/pArtour/networking-server/internal/database"
	"github.com/pArtour/networking-server/internal/models"
)

// UserService is a struct that contains a database connection
type UserService struct {
	db *database.Db
}

// NewUserService returns a new UserService struct
func NewUserService(db *database.Db) *UserService {
	return &UserService{
		db: db,
	}
}

// GetUsers returns all users
func (s *UserService) GetUsers() ([]*models.User, error) {
	rows, err := s.db.Conn.Query(context.Background(), "SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user *models.User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *UserService) GetUserByEmail(email string) (*models.UserWithPassword, error) {
	var user models.UserWithPassword
	err := s.db.Conn.QueryRow(context.Background(), "SELECT id, name, email, password FROM users WHERE email=$1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user
func (s *UserService) CreateUser(body *models.CreateUserInput) (*models.User, error) {
	var user models.User
	err := s.db.Conn.QueryRow(context.Background(), "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email", body.Name, body.Email, body.Password).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates a user
func (s *UserService) UpdateUser(id int64, body *models.UpdateUserInput) error {
	_, err := s.db.Conn.Exec(context.Background(), "UPDATE users SET name=$2 email=$3 WHERE id=$1", id, body.Name, body.Email)
	return err
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id int64) error {
	_, err := s.db.Conn.Exec(context.Background(), "DELETE FROM users WHERE id=$1", id)
	return err
}

// GetUser returns a user
func (s *UserService) GetUser(id int64) (*models.User, error) {
	var user models.User
	err := s.db.Conn.QueryRow(context.Background(), "SELECT id, name FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
