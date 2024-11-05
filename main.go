package main

import (
    // "audio_conversion/docs" 
    "audio_conversion/internal/router"

    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
)

// @title Audio Conversion API
// @version 1.0
// @description This is a simple audio conversion API.
// @host localhost:8080
// @BasePath /
func main() {
    r := gin.Default()

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    router.SetupRoutes(r)
    r.Run(":8080")
}
