package test

import (
    "testing"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/websocket/v2"

    "audio_conversion/internal/handler"
    "audio_conversion/internal/router"
)

func setupTestApp() *fiber.App {
    app := fiber.New()
    router.SetupRoutes(app)
    return app
}

func TestWebSocketConnection(t *testing.T) {
    app := setupTestApp()

    app.Get("/ws", websocket.New(handler.WebSocketConnection))
}
