package middleware

import (
	"context"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(authClient *auth.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
		}

		// Extract the ID token from the header
		idToken := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
		if idToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Authorization header"})
		}

		// Verify the ID token
		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
		}

		// Store the user ID in the context
		c.Locals("uid", token.UID)

		return c.Next()
	}
}
