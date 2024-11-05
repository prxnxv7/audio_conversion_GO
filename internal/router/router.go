package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"

	"audio_conversion/internal/handler"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/ws", websocket.New(handler.WebSocketConnection))
}
