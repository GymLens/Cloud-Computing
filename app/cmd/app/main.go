package main

import (
	"log"

	"github.com/GymLens/Cloud-Computing/config"
	"github.com/GymLens/Cloud-Computing/internal/server"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Relying on environment variables.")
	}

	cfg := config.LoadConfig()

	srv := server.NewServer(cfg)

	srv.App.Use(logger.New())

	srv.Start()
}
