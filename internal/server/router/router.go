package router

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/GymLens/Cloud-Computing/internal/server/controller"
	"github.com/GymLens/Cloud-Computing/internal/server/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, firebaseApp *firebase.App) {
	apiGroup := app.Group("/api")

	apiGroup.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "pong",
		})
	})

	authController := controller.NewAuthController(firebaseApp)
	userController := controller.NewUserController(firebaseApp)

	apiGroup.Post("/signup", authController.SignUp)
	apiGroup.Post("/signin", authController.SignIn)

	// Protected routes
	protected := apiGroup.Group("/user")
	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("Failed to initialize Firebase Auth client: %v", err)
	}
	protected.Use(middleware.AuthMiddleware(authClient))

	protected.Get("/profile", userController.GetProfile)
	protected.Post("/upload-avatar", userController.UploadAvatar)
}
