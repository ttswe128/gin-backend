package main

import (
	"log"

	"github.com/ttswe128/gin-backend/internal/config"
	"github.com/ttswe128/gin-backend/internal/routes"
)

func main() {
	cfg := config.Load()

	r := routes.SetupRouter(cfg.GinMode)

	log.Printf("Starting server on port %s in %s mode", cfg.Port, cfg.GinMode)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
