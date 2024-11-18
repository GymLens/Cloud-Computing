package auth

import (
	"context"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

type FirebaseAuth struct {
	Client *auth.Client
}

func NewFirebaseAuth() (*FirebaseAuth, error) {
	serviceAccountKeyPath := os.Getenv("FIREBASE_CONFIG")
	if serviceAccountKeyPath == "" {
		log.Println("FIREBASE_CONFIG is not set. Firebase authentication is disabled.")
		return nil, nil
	}

	opt := option.WithCredentialsFile(serviceAccountKeyPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("Failed to initialize Firebase App: %v", err)
		return nil, err
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Printf("Failed to get Firebase Auth client: %v", err)
		return nil, err
	}

	return &FirebaseAuth{Client: client}, nil
}

func (fa *FirebaseAuth) Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing Authorization header",
			})
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Authorization header format. Expected 'Bearer <token>'",
			})
		}

		idToken := parts[1]

		token, err := fa.Client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			log.Printf("Token verification failed: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		c.Locals("user", token)

		return c.Next()
	}
}
