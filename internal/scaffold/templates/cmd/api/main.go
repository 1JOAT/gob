package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/1joat/gob/internal/scaffold/templates/internal/config"
	"github.com/1joat/gob/internal/scaffold/templates/internal/database"
	"github.com/1joat/gob/internal/scaffold/templates/internal/health"
	"github.com/1joat/gob/internal/scaffold/templates/internal/middleware"
	"github.com/1joat/gob/internal/scaffold/templates/internal/routes"
)

func main() {
	cfg := config.Load()

	// Initialize Database
	_, err := database.Connect(cfg)
	if err != nil {
		log.Printf("Warning: Database connection failed: %v", err)
	}

	mux := http.NewServeMux()

	// Register routes
	health.Register(mux)
	routes.Register(mux)

	// Apply middleware
	handler := middleware.Logger(mux)

	fmt.Printf("Server starting on http://localhost%s\n", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, handler); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
