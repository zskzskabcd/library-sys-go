package router

import (
	"library-sys-go/docs"
	v1 "library-sys-go/internal/api/v1"
	"library-sys-go/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// cors
	r.Use(middleware.Cors())
	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

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
