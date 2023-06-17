package v1

import (
	"library-sys-go/internal/api"
	"library-sys-go/internal/middleware"
	"library-sys-go/internal/model"
	"library-sys-go/pkg/resp"
	"time"

	"github.com/gin-gonic/gin"
)

type SaveReservationReq struct {
	BookID uint `json:"bookId" binding:"required"`
	Retain int  `json:"retain" binding:"required"`
}

// 预约 godoc
// @Summary 预约
// @Description 预约
// @Tags 预约
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param req body SaveReservationReq true "预约信息"
// @Success 200 {object} resp.Resp
// @Router /reservation/save [post]
func SaveReservation(c *gin.Context) {
	var req SaveReservationReq
	user := c.MustGet("user").(*middleware.UserClaims)
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	// 检查该书的库存是否充足
	book := model.Book{}
	err := book.Query().Where("id = ?", req.BookID).First(&book).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	if book.Stock <= 0 {
		resp.Error(c, resp.CodeParamsInvalid, "该书库存不足")
		return
	}
	// 检查该读者是否已经预约过该书
	reservation := model.Reservation{}
	err = reservation.Query().Where("reader_id = ? AND book_id = ? AND status = ?", user.ID, req.BookID, model.ReservationStatusPending).First(&reservation).Error
	if err == nil {
		resp.Error(c, resp.CodeParamsInvalid, "已经预约过该书了")
		return
	}
	// 开启事务
	tx := model.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 预约
	reservation = model.Reservation{
		ReaderID: user.ID,
		BookID:   req.BookID,
		Status:   model.ReservationStatusPending,
	}
	err = reservation.Query().Save(&reservation).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	// 该书库存减一
	book.Stock--
	err = book.Query().Save(&book).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	tx.Commit()
	resp.Success(c)
}

// 取消预约 godoc
// @Summary 取消预约
// @Description 取消预约
// @Tags 预约
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id query int true "预约ID"
// @Success 200 {object} resp.Resp
// @Router /reservation/cancel [post]
func CancelReservation(c *gin.Context) {
	var req struct {
		ID int `json:"id" binding:"required" form:"id"`
	}
	user := c.MustGet("user").(*middleware.UserClaims)
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	reservation := model.Reservation{}
	err := reservation.Query().Where("id = ?", req.ID).Where("reader_id = ? AND status = ?", user.ID, model.ReservationStatusPending).First(&reservation).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, "该预约不存在")
		return
	}
	if reservation.Status != model.ReservationStatusPending {
		resp.Error(c, resp.CodeParamsInvalid, "该预约已取消或已完成")
		return
	}
	// 开启事务
	tx := model.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 取消预约
	reservation.Status = model.ReservationStatusCancel
	err = reservation.Query().Save(&reservation).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	// 该书库存加一
	book := model.Book{}
	err = book.Query().Where("id = ?", reservation.BookID).First(&book).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	book.Stock++
	err = book.Query().Save(&book).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	tx.Commit()
	resp.Success(c)
}

// 获取预约列表 godoc
// @Summary 获取预约列表
// @Description 获取预约列表
// @Tags 预约
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param readerId query int false "读者ID"
// @Param bookId query int false "书籍ID"
// @Param studentNo query string false "学号"
// @Param phone query string false "手机号"
// @Param readerName query string false "读者姓名"
// @Param bookName query string false "书籍名称"
// @Param from query string false "预约开始时间"
// @Param to query string false "预约结束时间"
// @Param status query int false "预约状态"
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} resp.RespList[model.Reservation]
// @Router /reservation/list [get]
func GetReservationList(c *gin.Context) {
	var req struct {
		ReaderID   int       `json:"readerId" form:"readerId"`
		BookID     int       `json:"bookId" form:"bookId"`
		StudentNo  string    `json:"studentNo" form:"studentNo"`
		Phone      string    `json:"phone" form:"phone"`
		ReaderName string    `json:"readerName" form:"readerName"`
		BookName   string    `json:"bookName" form:"bookName"`
		From       time.Time `json:"from" form:"from"`
		To         time.Time `json:"to" form:"to"`
		Status     uint8     `json:"status" form:"status"`
		api.Pagination
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	reservation := model.Reservation{}
	query := reservation.Query().Joins("LEFT JOIN readers ON readers.id = reservations.reader_id").Joins("LEFT JOIN books ON books.id = reservations.book_id")
	if req.ReaderID != 0 {
		query = query.Where("reader_id = ?", req.ReaderID)
	}
	if req.BookID != 0 {
		query = query.Where("book_id = ?", req.BookID)
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
		query = query.Where("created_at >= ?", req.From)
	}
	if !req.To.IsZero() {
		query = query.Where("created_at <= ?", req.To)
	}
	if req.Status != 0 {
		query = query.Where("status = ?", req.Status)
	}
	var total int64
	// 查询出删除的项目
	query.Unscoped().Preload("Reader", model.UnscopedQuery).Preload("Book", model.UnscopedQuery)
	var list []model.Reservation
	err := query.Count(&total).Offset(req.Offset()).Limit(req.Limit()).Order("id DESC").Find(&list).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	resp.SuccessList(c, list, total)
}

// 读者获取预约列表 godoc
// @Summary 读者获取预约列表
// @Description 读者获取预约列表
// @Tags 预约
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param status query int false "预约状态"
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} resp.RespList[model.Reservation]
// @Router /reservation/reader/list [get]
func GetReaderReservationList(c *gin.Context) {
	user := c.MustGet("user").(*middleware.UserClaims)
	var req struct {
		Status uint8 `json:"status" form:"status"`
		api.Pagination
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	reservations := []model.Reservation{}
	query := model.DB.Model(&model.Reservation{})
	query = query.Where("reader_id = ?", user.ID)
	if req.Status != 0 {
		query = query.Where("status = ?", req.Status)
	}
	var total int64
	query.Preload("Book", model.UnscopedQuery).Preload("Reader", model.UnscopedQuery)
	err := query.Count(&total).Offset(req.Offset()).Limit(req.Limit()).Order("id DESC").Find(&reservations).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	resp.SuccessList(c, reservations, total)
}
