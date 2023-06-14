package v1

import (
	"library-sys-go/internal/api"
	"library-sys-go/internal/model"
	"library-sys-go/pkg/douban"
	"library-sys-go/pkg/resp"
	"regexp"
	"strconv"
	"strings"

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
	book.OriginStock = book.Stock
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
// @Param force query bool false "强制删除"
// @Success 200 {object} resp.Resp
// @Router /book [delete]
func DeleteBook(c *gin.Context) {
	var req struct {
		ID    uint `json:"id" binding:"required" query:"id" form:"id"`
		Force bool `json:"force" query:"force" form:"force"`
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	book := model.Book{}
	book.ID = req.ID
	err := book.Query().Find(&book).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	if req.Force {
		err = book.Query().Delete(&book).Error
		if err != nil {
			resp.Error(c, resp.CodeInternalServer, err.Error())
			return
		}
		resp.Success(c)
	}
	// 检查剩余库存
	if book.Stock != book.OriginStock && !req.Force {
		resp.Error(c, resp.CodeParamsInvalid, "有未还的书籍，不能删除")
		return
	}
	// 软删除
	err = book.Query().Delete(&book).Error
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
		Keyword string `json:"keyword" form:"keyword"`
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
	err = booksQuery.Offset(req.Offset()).Limit(req.Size).Find(&books).Error
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
		ID int `json:"id" binding:"required" form:"id"`
	}
	// type res resp.Resp[model.Book]
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

// 云搜索书籍 godoc
// @Summary 云搜索书籍
// @Description 云搜索书籍
// @Tags 书籍
// @Accept json
// @Produce json
// @Param keyword query string true "关键字"
// @Param start query int false "起始位置"
// @Param count query int false "数量"
// @Success 200 {object} resp.RespList[model.Book]
// @Router /book/search [get]
func SearchBook(c *gin.Context) {
	req := struct {
		Keyword string `json:"keyword" form:"keyword"`
		Start   int    `json:"start" form:"start"`
		Count   int    `json:"count" form:"count"`
	}{
		Start: 0,
		Count: 10,
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	result, err := douban.DouBanSearch(req.Keyword, req.Start, req.Count)
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	books := make([]model.Book, 0)

	for _, item := range result.Books {
		// 正则匹配价格 38.00元
		price_str := regexp.MustCompile(`\d+\.\d+`).FindString(item.Price)
		price, _ := strconv.ParseFloat(price_str, 64)
		books = append(books, model.Book{
			Title:       item.Title,
			Author:      strings.Join(item.Author, ", "),
			ISBN:        item.Isbn13,
			Publisher:   item.Publisher,
			PublishDate: item.Pubdate,
			Summary:     item.Summary,
			Cover:       item.Image,
			Price:       price,
		})
	}
	resp.SuccessList(c, books, int64(result.Total))
}
