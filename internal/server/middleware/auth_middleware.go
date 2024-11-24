package middleware

import (
	"context"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(authClient *auth.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
		}

		idToken := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
		if idToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Authorization header"})
		}

		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
		}

		c.Locals("uid", token.UID)

		return c.Next()
	}
}
