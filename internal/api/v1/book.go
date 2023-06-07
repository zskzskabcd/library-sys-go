package v1

import (
	"library-sys-go/internal/api"
	"library-sys-go/internal/model"
	"library-sys-go/pkg/resp"

	"github.com/gin-gonic/gin"
)

// 新书入库 | 更新书籍信息
// Request Body: {title, author, publisher, publish_date, summary, price, stock}
func SaveBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	err := book.Query().Save(&book).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	resp.Success(c)
}

// 删除书籍
func DeleteBook(c *gin.Context) {
	var req struct {
		ID int `json:"id" binding:"required"`
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	book := model.Book{}
	err := book.Query().Delete(&book).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	resp.Success(c)
}

// 查询书籍列表
func ListBook(c *gin.Context) {
	var req struct {
		Keyword string `json:"keyword"`
		api.Pagination
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	book := model.Book{}
	booksQuery := book.Query()
	if req.Keyword != "" {
		booksQuery = booksQuery.Where("title LIKE ?", "%"+req.Keyword+"%")
	}
	var books []model.Book
	var total int64
	err := booksQuery.Count(&total).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	err = booksQuery.Offset((req.Page - 1) * req.Size).Limit(req.Size).Find(&books).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	resp.SuccessList(c, books, total)
}

// 查询书籍详情
func GetBook(c *gin.Context) {
	var req struct {
		ID int `json:"id" binding:"required"`
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	book := model.Book{}
	err := book.Query().Where("id = ?", req.ID).First(&book).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	resp.SuccessData(c, book)
}
