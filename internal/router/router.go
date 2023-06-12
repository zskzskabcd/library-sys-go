package router

import (
	v1 "library-sys-go/internal/api/v1"
	"library-sys-go/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zskzskabcd/knife4g"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// cors
	r.Use(middleware.Cors())
	// swagger
	r.GET("/doc/*any", knife4g.Handler(knife4g.Config{RelativePath: "/doc"}))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.URL("/doc/front/docJson")))
	// Common
	api := r.Group("/api/v1")
	{
		// ping
		api.GET("/ping", v1.Ping)
	}
	// 匿名接口
	{
		// 读者登陆
		api.POST("/reader/login", v1.LoginReader)
		// 管理员登陆
		api.POST("/admin/login", v1.AdminLogin)
	}
	// 需要登陆
	api.Use(middleware.LoginAuthMiddleware())
	reader := api.Use(middleware.ReaderAuthMiddleware())
	// Reader
	{
		// 读者修改密码
		reader.POST("/reader/password", v1.UpdateReaderPassword)
		// 借书
		reader.POST("/lending/book", v1.LendBook)
		// 还书
		reader.POST("/return/book", v1.ReturnBook)
		// 查询借阅记录
		reader.GET("/lending/listByReader", v1.ListLendingByReader)
		// 查询借阅记录详情
		reader.GET("/lend/record/:id", v1.LendingDetail)
		// 预约
		reader.POST("/reservation/save", v1.SaveReservation)
		// 取消预约
		reader.POST("/reservation/cancel", v1.CancelReservation)
		// 查询预约记录 读者
		reader.GET("/reservation/reader/list", v1.GetReaderReservationList)
		// 查询书籍详情
		reader.GET("/book/get", v1.GetBook)
		// 查询书籍列表
		reader.GET("/book/list", v1.ListBook)
	}
	// 管理员
	admin := api.Use(middleware.AdminAuthMiddleware())
	{
		// 管理员修改密码
		admin.POST("/admin/password", v1.AdminChangePassword)
		// 新书入库 & 修改书籍信息
		admin.POST("/book", v1.SaveBook)
		// 删除书籍
		admin.DELETE("/book", v1.DeleteBook)
		// 新增读者
		admin.POST("/reader", v1.SaveReader)
		// 删除读者
		admin.DELETE("/reader", v1.DeleteReader)
		// 查询读者列表
		admin.GET("/reader/list", v1.ListReader)
		// 查询读者详情
		admin.GET("/reader", v1.GetReader)
		// 查询借阅记录
		admin.GET("/lending/list", v1.ListLending)
		// 查询借阅记录详情
		admin.GET("/lending/detail", v1.LendingDetail)
		// 查询预约记录
		admin.GET("/reservation/list", v1.GetReservationList)
	}

	return r
}
