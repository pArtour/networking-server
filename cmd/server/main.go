package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/pArtour/networking-server/internal/config"
	"github.com/pArtour/networking-server/internal/controllers"
	"github.com/pArtour/networking-server/internal/database"
	"github.com/pArtour/networking-server/internal/errors"
	"github.com/pArtour/networking-server/internal/handlers"
	"github.com/pArtour/networking-server/internal/services"
	"log"
)

func main() {
	cfg := config.NewConfig()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return ctx.Status(code).JSON(&errors.ErrorResponse{Code: code, Message: err.Error()})
		},
	})
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return cfg.Env == "development"
		},
	}))

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

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
