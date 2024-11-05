package main

import (
    "audio_conversion/internal/router"

    "github.com/gin-gonic/gin"
)

// @title Audio Conversion API
// @version 1.0
// @description This is a simple audio conversion API.
// @host localhost:8080
// @BasePath /
func main() {
    r := gin.Default()
    router.SetupRoutes(r)
    r.Run(":8080")
}
