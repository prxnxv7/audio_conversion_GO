package handler

import (
    "audio_conversion/internal/audio"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "net/http"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func WebSocketConnection(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        return
    }
    defer conn.Close()

    for {
        _, wavData, err := conn.ReadMessage()
        if err != nil {
            break
        }

        flacData, err := audio.ConvertWAVToFLAC(wavData)
        if err != nil {
            continue
        }

        err = conn.WriteMessage(websocket.BinaryMessage, flacData)
        if err != nil {
            break
        }
    }
}
