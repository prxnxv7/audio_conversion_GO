package router

import (
	"audio_conversion/internal/handler"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
   	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
    r.GET("/ws", handler.WebSocketConnection)
}
