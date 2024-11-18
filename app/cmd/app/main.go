package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // Health check route
    app.Get("/api/ping", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "message": "pong",
        })
    })

    // Start server
    port := "8080"
    log.Printf("Starting server on port %s", port)
    if err := app.Listen(":" + port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
