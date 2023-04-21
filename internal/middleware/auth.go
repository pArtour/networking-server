package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v3"
	"github.com/pArtour/networking-server/internal/config"
	"github.com/pArtour/networking-server/internal/errors"
)

func JWTProtected() fiber.Handler {
	return jwt.New(jwt.Config{
		SigningKey:   []byte(config.Cfg.JWTSecret),
		ErrorHandler: jwtErrorHandler,
	})
}

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(&errors.ErrorResponse{Code: fiber.StatusUnauthorized, Message: "Unauthorized"})
}
