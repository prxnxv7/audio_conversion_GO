package test

import (
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "audio_conversion/internal/router"
    "net/http/httptest"
)

func setupTestApp() *gin.Engine {
    app := gin.Default()
    router.SetupRoutes(app)
    return app
}

func TestWebSocketConnection(t *testing.T) {
    app := setupTestApp()

    ts := httptest.NewServer(app)
    defer ts.Close()

    url := "ws" + ts.URL[4:] + "/ws" 
    conn, _, err := websocket.DefaultDialer.Dial(url, nil)
    if err != nil {
        t.Fatalf("Failed to connect to WebSocket: %v", err)
    }
    defer conn.Close()

}
