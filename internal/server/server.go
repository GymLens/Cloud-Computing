package server

import (
	"log"

	"github.com/GymLens/Cloud-Computing/config"
	"github.com/GymLens/Cloud-Computing/internal/server/router"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App  *fiber.App
	Port string
}

func NewServer(cfg *config.Config) *Server {
	app := fiber.New()

	router.SetupRoutes(app)

	return &Server{
		App:  app,
		Port: cfg.Port,
	}
}

func (s *Server) Start() {
	log.Printf("Starting server on port %s", s.Port)
	if err := s.App.Listen(":" + s.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
