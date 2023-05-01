package services

import (
	"context"
	"github.com/pArtour/networking-server/internal/database"
	"github.com/pArtour/networking-server/internal/models"
)

type ConnectionService struct {
	Db *database.Db
}

func NewConnectionService(db *database.Db) *ConnectionService {
	return &ConnectionService{
		Db: db,
	}
}

// GetConnectionsByUserId returns all connections of a user
func (s *ConnectionService) GetConnectionsByUserId(userId int64) ([]models.Connection, error) {
	rows, err := s.Db.Conn.Query(context.Background(), "SELECT c.id, c.user_id_1, c.user_id_2 FROM connections c WHERE c.user_id_1 = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var connections []models.Connection
	for rows.Next() {
		var connection models.Connection
		err := rows.Scan(&connection.Id, &connection.UserId1, &connection.UserId2)
		if err != nil {
			return nil, err
		}
		connections = append(connections, connection)
	}

	if connections == nil {
		return []models.Connection{}, nil
	}

	return connections, nil
}

// CreateConnection creates a new connection.
func (s *ConnectionService) CreateConnection(body *models.CreateConnectionRecordInput) (*models.Connection, error) {
	var connection models.Connection
	err := s.Db.Conn.QueryRow(context.Background(), "INSERT INTO connections (user_id_1, user_id_2) VALUES ($1, $2) RETURNING id, user_id_1, user_id_2", body.UserId, body.TargetUserId).Scan(&connection.Id, &connection.UserId1, &connection.UserId2)

	if err != nil {
		return nil, err
	}
	return &connection, nil
}

// DeleteConnection deletes a connection
func (s *ConnectionService) DeleteConnection(id int64) error {
	_, err := s.Db.Conn.Exec(context.Background(), "DELETE FROM connections WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// GetConnectionById returns a connection by id
func (s *ConnectionService) GetConnectionById(id int64) (*models.Connection, error) {
	var connection models.Connection
	err := s.Db.Conn.QueryRow(context.Background(), "SELECT id, user_id_1, user_id_2 FROM connections WHERE id = $1", id).Scan(&connection.Id, &connection.UserId1, &connection.UserId2)
	if err != nil {
		return nil, err
	}
	return &connection, nil
}

// DeleteUserConnection deletes a connection by user id
func (s *ConnectionService) DeleteUserConnection(userId int64, connecton_id int64) error {
	_, err := s.Db.Conn.Exec(context.Background(), "DELETE FROM user_connections WHERE user_id = $1 AND connection_id = $2", userId, connecton_id)
	if err != nil {
		return err
	}
	return nil
}
