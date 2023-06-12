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
	// 检查启动参数是否有init
	// Args := os.Args
	// if len(Args) > 0 && Args[0] == "init" {
	// 	v1.AddAdmin()
	// 	return
	// }

	// swagger文档
	// @title 图书管理系统API文档
	// @description 图书管理系统API文档
	// @contact.name wear工程师
	// @basePath /api/v1
	// @version v1

	model.Setup()
	server := router.InitRouter()
	log.Println("server run at 8761")
	log.Println("访问 http://127.0.0.1:8761/doc/index 获取swagger文档")
	err := server.Run(":8761")
	if err != nil {
		return
	}
}
