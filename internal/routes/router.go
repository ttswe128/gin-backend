package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ttswe128/gin-backend/internal/handlers"
)

func SetupRouter(ginMode string) *gin.Engine {
	gin.SetMode(ginMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", handlers.HealthCheck)
	}

	return r
}
