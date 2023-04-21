package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

var conn *pgx.Conn

func setupDatabase() {
	var err error
	conn, err = pgx.Connect(context.Background(), "postgresql://postgres:NZ4ZGECDSJWE9JwnFnlN@containers-us-west-136.railway.app:7316/railway")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
}
