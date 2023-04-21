package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/pArtour/networking-server/internal/config"
)

type Db struct {
	Conn *pgx.Conn
}

func NewDb(config *config.Config) *Db {
	var err error
	conn, err := pgx.Connect(context.Background(), config.Database.Url)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return nil
	}

	return &Db{
		Conn: conn,
	}
}
