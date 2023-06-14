package v1

import (
	"errors"
	"library-sys-go/internal/api"
	"library-sys-go/internal/middleware"
	"library-sys-go/internal/model"
	"library-sys-go/pkg/resp"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 借阅相关

// 借书 godoc
// @Summary 借书
// @Description 借书
// @Tags 借阅
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param bookId query int true "书籍ID"
// @Param days query int true "借阅时长（天）"
// @Success 200 {object} resp.Resp
// @Router /lending/book [post]
func LendBook(c *gin.Context) {
	var req struct {
		BookID uint `json:"bookId" form:"bookId" binding:"required"`
		// 借阅时长（天）
		Days int `json:"days" form:"days" binding:"required"`
	}
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	user := c.MustGet("user").(*middleware.UserClaims)
	// 查询书籍是否存在
	book := model.Book{}
	err := book.Query().Where("id = ?", req.BookID).First(&book).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	// 查询读者是否存在
	reader := model.Reader{}
	err = reader.Query().Where("id = ?", user.ID).First(&reader).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	// 查询书籍库存
	if book.Stock <= 0 {
		resp.Error(c, resp.CodeParamsInvalid, "库存不足")
		return
	}
	// 查询是否已经借阅
	lending := model.Lending{}
	err = lending.Query().Where("book_id = ? AND reader_id = ?", req.BookID, user.ID).First(&lending).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	if lending.ID != 0 {
		resp.Error(c, resp.CodeParamsInvalid, "你已经借阅该书籍")
		return
	}
	// 借阅
	// 开启事务
	tx := model.DB.Begin()
	// 书籍库存减一
	err = tx.Model(&book).Update("stock", book.Stock-1).Error
	if err != nil {
		tx.Rollback()
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	// 生成借阅记录
	lending = model.Lending{
		BookID:    req.BookID,
		ReaderID:  user.ID,
		LendTime:  time.Now(),
		ReturnTim: time.Now().AddDate(0, 0, req.Days),
		Status:    model.LendingStatusLending,
	}
	err = tx.Create(&lending).Error
	if err != nil {
		tx.Rollback()
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	// 查询是否有预约
	reservation := model.Reservation{}
	reservation.Query().Where("book_id = ? AND reader_id = ?", req.BookID, user.ID).First(&reservation)
	if reservation.ID != 0 {
		// 有预约，更新预约状态
		reservation.Status = model.ReservationStatusSuccess
		err = tx.Save(&reservation).Error
		if err != nil {
			tx.Rollback()
			resp.Error(c, resp.CodeInternalServer, err.Error())
			return
		}
		// 书籍库存加一
		err = tx.Model(&book).Update("stock", book.Stock+1).Error
		if err != nil {
			tx.Rollback()
			resp.Error(c, resp.CodeInternalServer, err.Error())
			return
		}
	}
	// 提交事务
	tx.Commit()
	resp.Success(c)
}

// 还书 godoc
// @Summary 还书
// @Description 还书
// @Tags 借阅
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param bookId query int true "书籍ID"
// @Success 200 {object} resp.Resp
// @Router /return/book [post]
func ReturnBook(c *gin.Context) {
	var req struct {
		BookID int `json:"bookId" binding:"required" form:"bookId"`
	}
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	user := c.MustGet("user").(*middleware.UserClaims)
	// 查询书籍是否存在
	book := model.Book{}
	err := book.Query().Where("id = ?", req.BookID).First(&book).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	// 查询读者是否存在
	reader := model.Reader{}
	err = reader.Query().Where("id = ?", user.ID).First(&reader).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	// 查询是否已经借阅
	lending := model.Lending{}
	err = lending.Query().Where("book_id = ? AND reader_id = ?", req.BookID, user.ID).First(&lending).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	if lending.ID == 0 {
		resp.Error(c, resp.CodeParamsInvalid, "你没有借阅该书籍")
		return
	}
	// 还书
	// 开启事务
	tx := model.DB.Begin()
	// 书籍库存加一
	err = tx.Model(&book).Update("stock", book.Stock+1).Error
	if err != nil {
		tx.Rollback()
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	// 更新借阅记录
	err = tx.Model(&lending).Update("status", model.LendingStatusReturned).Error
	if err != nil {
		tx.Rollback()
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	tx.Commit()
	resp.Success(c)
}

// 借阅记录
// @Summary 借阅记录
// @Description 借阅记录
// @Tags 借阅
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param readerId query int false "读者ID"
// @Param bookId query int false "书籍ID"
// @Param studentNo query string false "学号"
// @Param phone query string false "手机号"
// @Param readerName query string false "读者姓名"
// @Param bookName query string false "书籍名称"
// @Param from query string false "借阅时间开始"
// @Param to query string false "借阅时间结束"
// @Param status query int false "状态"
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} resp.Resp{data=[]model.Lending}
// @Router /lending/list [get]
func ListLending(c *gin.Context) {
	var req struct {
		ReaderID   int       `json:"readerId" form:"readerId" example:"1"`
		BookID     int       `json:"bookId" form:"bookId" example:"1"`
		StudentNo  string    `json:"studentNo" form:"studentNo" example:"201800000000"`
		Phone      string    `json:"phone" form:"phone" example:"18888888888"`
		ReaderName string    `json:"readerName" form:"readerName" example:"张三"`
		BookName   string    `json:"bookName" form:"bookName" example:"三国演义"`
		From       time.Time `json:"from" form:"from"`
		To         time.Time `json:"to" form:"to"`
		Status     uint8     `json:"status" form:"status" example:"1"`
		api.Pagination
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	lendings := make([]model.Lending, 0)
	query := model.DB.Model(&model.Lending{}).Joins("LEFT JOIN readers ON readers.id = lendings.reader_id").Joins("LEFT JOIN books ON books.id = lendings.book_id")
	if req.ReaderID != 0 {
		query = query.Where("lendings.reader_id = ?", req.ReaderID)
	}
	if req.BookID != 0 {
		query = query.Where("lendings.book_id = ?", req.BookID)
	}
	if req.StudentNo != "" {
		query = query.Where("readers.student_no = ?", req.StudentNo)
	}
	if req.Phone != "" {
		query = query.Where("readers.phone = ?", req.Phone)
	}
	if req.ReaderName != "" {
		query = query.Where("readers.name = ?", req.ReaderName)
	}
	if req.BookName != "" {
		query = query.Where("books.name = ?", req.BookName)
	}
	if !req.From.IsZero() {
		query = query.Where("lendings.lend_time >= ?", req.From)
	}
	if !req.To.IsZero() {
		query = query.Where("lendings.lend_time <= ?", req.To)
	}
	if req.Status != 0 {
		query = query.Where("lendings.status = ?", req.Status)
	}
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	query.Preload("Reader").Preload("Book")
	err = query.Offset(req.Offset()).Limit(req.Limit()).Find(&lendings).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	resp.SuccessList(c, lendings, total)
}

// LendingDetail godoc
// @Summary 借阅记录详情
// @Description 借阅记录详情
// @Tags 借阅
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param lendingId query int true "借阅记录ID"
// @Success 200 {object} resp.Resp{data=model.Lending}
// @Router /lending/detail [get]
func LendingDetail(c *gin.Context) {
	var req struct {
		LendingID int `json:"lendingId" form:"lendingId" binding:"required"`
	}
	// 用户只能查询自己的借阅记录
	user := c.MustGet("user").(*middleware.UserClaims)
	var readerID uint
	if !middleware.RoleContains(user.Role, middleware.RoleAdmin) {
		readerID = user.ID
	}
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	lending := model.Lending{}
	query := model.DB.Model(&model.Lending{}).Where("id = ?", req.LendingID)
	if readerID != 0 {
		query = query.Where("reader_id = ?", readerID)
	}
	err := query.Preload("Book").Preload("Reader").First(&lending).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	resp.SuccessData(c, lending)
}

// LendingCreate godoc
// @Summary 读者查询借阅记录
// @Description 读者查询借阅记录
// @Tags 借阅
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param bookId query int true "书籍ID"
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} resp.RespList[model.Lending]
// @Router /lending/listByReader [get]
func ListLendingByReader(c *gin.Context) {
	var req struct {
		BookID int `json:"bookId" form:"bookId" example:"1"` // 书籍ID
		api.Pagination
	}
	user := c.MustGet("user").(*middleware.UserClaims)
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	lendings := make([]model.Lending, 0)
	query := model.DB.Model(&model.Lending{}).Where("reader_id = ?", user.ID)
	if req.BookID != 0 {
		query = query.Where("book_id = ?", req.BookID)
	}
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	query.Preload("Reader").Preload("Book")
	err = query.Offset(req.Offset()).Limit(req.Limit()).Find(&lendings).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	resp.SuccessList(c, lendings, total)
}
