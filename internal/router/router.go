package router

import (
	v1 "gin-mini-starter/internal/api/v1"
	"gin-mini-starter/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// cors
	r.Use(middleware.Cors())

	// Common
	api := r.Group("/api")
	{
		// ping
		api.GET("/ping", v1.Ping)
	}
	// Auth
	api.Use(v1.Auth())
	{

	}

	return r
}
