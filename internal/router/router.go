package router

import (
	v1 "library-sys-go/internal/api/v1"
	"library-sys-go/internal/middleware"

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
	// 需要登陆
	api.Use(middleware.LoginAuthMiddleware())
	// Reader
	{

	}
	// 管理员

	return r
}
