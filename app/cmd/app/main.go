package main

import (
	"github.com/GymLens/Cloud-Computing/config"
	"github.com/GymLens/Cloud-Computing/internal/server"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the server with the loaded configuration
	srv := server.NewServer(cfg)

	// Start the server
	srv.Start()
}
