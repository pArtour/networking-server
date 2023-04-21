package helpers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pArtour/networking-server/internal/config"
	"time"
)

// GenerateJWTToken generates a JWT token
func GenerateJWTToken(userID int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString([]byte(config.Cfg.JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ExtractUserIDFromJWT(c *fiber.Ctx) (int64, error) {
	user := c.Locals("user").(*jwt.Token)

	if user == nil {
		return 0, errors.New("cannot extract user from JWT")
	}

	claims := user.Claims.(jwt.MapClaims)
	userID := int64(claims["user_id"].(float64))
	return userID, nil
}
