package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pArtour/networking-server/internal/config"
)

// Db is a struct that contains a database connection
type Db struct {
	Conn *pgxpool.Pool
}

// NewDb returns a new Db struct
func NewDb() *Db {
	var err error

	conn, err := pgxpool.Connect(context.Background(), config.Cfg.Database.Url)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return nil
	}

	return &Db{
		Conn: conn,
	}
}
