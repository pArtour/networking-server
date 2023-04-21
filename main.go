package main

import (
	"os"
	"strconv"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// Database setup
	setupDatabase()

	// Routes
	app.Get("/users", func(c *fiber.Ctx) error {
		users, err := getUsers()
		if err != nil {
			return c.Status(500).SendString("Error fetching users")
		}
		return c.JSON(users)
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		if name == "" {
			return c.Status(400).SendString("Name is required")
		}
		id, err := createUser(name)
		if err != nil {
			return c.Status(500).SendString("Error creating user")
		}
		return c.JSON(fiber.Map{"id": id})
	})

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(400).SendString("Invalid user ID")
		}
		name := c.FormValue("name")
		if name == "" {
			return c.Status(400).SendString("Name is required")
		}
		err = updateUser(id, name)
		if err != nil {
			return c.Status(500).SendString("Error updating user")
		}
		return c.SendStatus(200)
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(400).SendString("Invalid user ID")
		}
		err = deleteUser(id)
		if err != nil {
			return c.Status(500).SendString("Error deleting user")
		}
		return c.SendStatus(200)
	})

	port := getPort()
	// Start the server
	app.Listen(port)
}
