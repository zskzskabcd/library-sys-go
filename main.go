package main

import (
	"library-sys-go/internal/model"
	"library-sys-go/internal/router"
	"log"
)

func init() {
	//gin.SetMode(gin.ReleaseMode)
}

func main() {
	model.Setup()
	server := router.InitRouter()
	log.Println("server run at 8761")
	err := server.Run(":8761")
	if err != nil {
		return
	}
}
