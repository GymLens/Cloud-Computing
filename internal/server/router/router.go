package router

import (
	"github.com/GymLens/Cloud-Computing/internal/server/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Health check
	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "pong",
		})
	})

	// Initialize UserController
	userController := controller.NewUserController()

	// User routes
	api.Post("/users", userController.CreateUser)       // Create a new user
	api.Get("/users", userController.ListUsers)         // List all users
	api.Get("/users/:id", userController.GetUser)       // Get a user by ID
	api.Put("/users/:id", userController.UpdateUser)    // Update a user by ID
	api.Delete("/users/:id", userController.DeleteUser) // Delete a user by ID
}
