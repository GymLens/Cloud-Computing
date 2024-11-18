package router

import (
	"github.com/GymLens/Cloud-Computing/config"
	"github.com/GymLens/Cloud-Computing/internal/server/controller"
	"github.com/GymLens/Cloud-Computing/pkg/auth"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, firebaseAuth *auth.FirebaseAuth, cfg *config.Config, v *validator.Validate) {
	api := app.Group("/api")

	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "pong",
		})
	})

	authController := controller.NewAuthController(firebaseAuth, cfg, v)

	// Auth routes
	api.Post("/signup", authController.SignUp)
	api.Post("/signin", authController.SignIn)

	userController := controller.NewUserController()

	// User routes
	api.Post("/users", userController.CreateUser) // Create a new user
	api.Get("/users", userController.ListUsers)   // List all users
	api.Get("/users/:id", userController.GetUser) // Get a user by ID

	// Protected routes (Update and Delete)
	if firebaseAuth != nil && firebaseAuth.Client != nil {
		protected := api.Group("/users")
		protected.Use(firebaseAuth.Middleware()) // Apply Firebase Auth middleware

		protected.Put("/:id", userController.UpdateUser)    // Update a user by ID
		protected.Delete("/:id", userController.DeleteUser) // Delete a user by ID
	} else {
		// Log warning and block access if FirebaseAuth not init
		protected := api.Group("/users")
		protected.Use(func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Authentication is not enabled. Protected routes are unavailable.",
			})
		})

		protected.Put("/:id", func(c *fiber.Ctx) error {
			return nil
		})

		protected.Delete("/:id", func(c *fiber.Ctx) error {
			return nil
		})
	}
}
