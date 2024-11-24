package server

import (
	"fmt"
	"log"

	"github.com/GymLens/Cloud-Computing/config"
	"github.com/GymLens/Cloud-Computing/internal/server/router"
	"github.com/GymLens/Cloud-Computing/pkg/auth"
	"github.com/gofiber/fiber/v2"
)

func Start() error {
	app := fiber.New()
	cfg := config.GetConfig()

	// Firebase Init
	firebaseApp, err := auth.InitializeFirebaseApp()
	if err != nil {
		return err
	}
	if firebaseApp == nil {
		return fmt.Errorf("firebaseApp is nil")
	}

	router.SetupRoutes(app, firebaseApp)

	log.Printf("Server is running on port %s", cfg.Port)
	return app.Listen(":" + cfg.Port)
}
