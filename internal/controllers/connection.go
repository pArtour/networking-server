package controllers

import (
	"github.com/pArtour/networking-server/internal/models"
	"github.com/pArtour/networking-server/internal/services"
)

type ConnectionController struct {
	connectionService *services.ConnectionService
}

func NewConnectionController(cs *services.ConnectionService) *ConnectionController {
	return &ConnectionController{
		connectionService: cs,
	}
}

// GetUserConnections returns all connections for a user
func (c *ConnectionController) GetUserConnections(id int64) ([]models.Connection, error) {
	connections, err := c.connectionService.GetConnectionsByUserId(id)
	if err != nil {
		return nil, err
	}
	return connections, nil
}

// CreateConnection creates a new connection
func (c *ConnectionController) CreateConnection(body *models.CreateConnectionRecordInput) (*models.Connection, error) {
	connection, err := c.connectionService.CreateConnection(body)
	if err != nil {
		return nil, err
	}
	return connection, nil
}

// DeleteConnection deletes a connection
func (c *ConnectionController) DeleteConnection(id int64) error {
	err := c.connectionService.DeleteConnection(id)
	if err != nil {
		return err
	}
	return nil
}

// GetConnectionById returns a connection by id
func (c *ConnectionController) GetConnectionById(id int64) (*models.Connection, error) {
	connection, err := c.connectionService.GetConnectionById(id)
	if err != nil {
		return nil, err
	}
	return connection, nil
}

// DeleteUserConnection deletes all connections for a user
func (c *ConnectionController) DeleteUserConnection(userId, connectionId int64) error {
	err := c.connectionService.DeleteUserConnection(userId, connectionId)
	if err != nil {
		return err
	}
	return nil
}
