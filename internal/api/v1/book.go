package v1

import (
	"library-sys-go/internal/api"
	"library-sys-go/internal/model"
	"library-sys-go/pkg/resp"

	"github.com/gin-gonic/gin"
)

// 新书入库 | 更新书籍信息 godoc
// @Summary 新书入库 | 更新书籍信息
// @Description 新书入库 | 更新书籍信息
// @Tags 书籍
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param book body model.Book true "book"
// @Success 200 {object} resp.Resp
// @Router /book [post]
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

// 删除书籍 godoc
// @Summary 删除书籍
// @Description 删除书籍
// @Tags 书籍
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id query int true "id"
// @Success 200 {object} resp.Resp
// @Router /book [delete]
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

// 查询书籍列表 godoc
// @Summary 查询书籍列表
// @Description 查询书籍列表
// @Tags 书籍
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param keyword query string false "关键字"
// @Param page query int false "页码"
// @Param size query int false "每页条数"
// @Success 200 {object} resp.Resp{data=[]model.Book}
// @Router /book/list [get]
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

// 查询书籍详情 godoc
// @Summary 查询书籍详情
// @Description 查询书籍详情
// @Tags 书籍
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id query int true "id"
// @Success 200 {object} resp.Resp{data=model.Book}
// @Router /book/get [get]
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
