package main

import (
	"gin-mini-starter/model"
	"gin-mini-starter/router"
	"gin-mini-starter/utils/log"
)

func init() {
	//gin.SetMode(gin.ReleaseMode)
}

func main() {
	model.Setup()
	server := router.InitRouter()
	log.Debug("server running at 0.0.0.0:8899")
	err := server.Run(":8899")
	if err != nil {
		return
	}
}
