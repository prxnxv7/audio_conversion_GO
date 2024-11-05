package handler

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/websocket/v2"
    "audio_conversion/internal/audio"
)

func WebSocketHandler(c *fiber.Ctx) error {
    if websocket.IsWebSocketUpgrade(c) {
        c.Locals("allowed", true)
        return c.Next()
    }
    return fiber.ErrUpgradeRequired
}

func WebSocketConnection(c *websocket.Conn) {
    for {
        messageType, data, err := c.ReadMessage()
        if err != nil {
            log.Println("Error reading message:", err)
            break
        }

        if messageType == websocket.BinaryMessage {
            // Convert WAV data to FLAC
            flacData, err := audio.ConvertWAVToFLAC(data)
            if err != nil {
                log.Println("Error converting WAV to FLAC:", err)
                break
            }

            // Send the FLAC data back via WebSocket
            if err := c.WriteMessage(websocket.BinaryMessage, flacData); err != nil {
                log.Println("Error writing message:", err)
                break
            }
        }
    }
}
