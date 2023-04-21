package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/pArtour/networking-server/internal/config"
)

// Db is a struct that contains a database connection
type Db struct {
	Conn *pgx.Conn
}

// NewDb returns a new Db struct
func NewDb() *Db {
	var err error
	conn, err := pgx.Connect(context.Background(), config.Cfg.Database.Url)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return nil
	}

	return &Db{
		Conn: conn,
	}
}
