package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pArtour/networking-server/internal/config"
	"github.com/pArtour/networking-server/internal/controllers"
	"github.com/pArtour/networking-server/internal/database"
	"github.com/pArtour/networking-server/internal/handlers"
	"github.com/pArtour/networking-server/internal/services"
	"log"
)

func main() {
	cfg := config.NewConfig()

	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// Database setup
	db := database.NewDb(cfg)
	defer db.Conn.Close(context.Background())

	// Services setup
	s := services.NewServices(db)

	// Controllers setup
	c := controllers.NewControllers(s)

	// Routes setup
	r := app.Group("/api/v1")

	// Handlers setup
	handlers.NewHandlers(r, c)

	// Start the server
	log.Fatal(app.Listen(cfg.Server.Port))
}
