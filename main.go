package main

import (
    "log"
    
    "github.com/gofiber/fiber/v2"

    "audio_conversion/config"
    "audio_conversion/internal/router"
)

func main() {
    cfg := config.LoadConfig()
    app := fiber.New()
    router.SetupRoutes(app)

    log.Printf("Starting server on port %s...", cfg.Port)
    if err := app.Listen(":" + cfg.Port); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
