package server

import (
	"log"

	"github.com/GymLens/Cloud-Computing/config"
	"github.com/GymLens/Cloud-Computing/internal/server/router"
	"github.com/GymLens/Cloud-Computing/pkg/auth"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App  *fiber.App
	Port string
}

func NewServer(cfg *config.Config) *Server {
	app := fiber.New()

	v := validator.New()

	firebaseAuth, err := auth.NewFirebaseAuth()
	if err != nil {
		log.Fatalf("Failed to initialize Firebase Auth: %v", err)
	}

	router.SetupRoutes(app, firebaseAuth, cfg, v)

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
