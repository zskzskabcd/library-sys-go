package model

import "gorm.io/gorm"

// 预约表（reservation）
// id (int，主键)
// book_id (int，外键，关联书籍表)
// reader_id (int，外键，关联读者表)

type Reservation struct {
	Model
	BookID   uint              `json:"bookId" gorm:"type:int;not null" binding:"required" example:"1"` // 书籍ID
	Book     Book              `json:"book" gorm:"foreignkey:book_id"`
	ReaderID uint              `json:"readerId" gorm:"type:int;not null" binding:"required" example:"1" ` // 读者ID
	Reader   Reader            `json:"reader" gorm:"foreignkey:reader_id"`
	Status   ReservationStatus `json:"status" gorm:"type:tinyint;not null" binding:"required" example:"1" enums:"1,2,3,4" enumsdes:"1:进行中,2:已借阅,3:已超时,4:已取消"` // 预约状态
}

func (r *Reservation) Query() *gorm.DB {
	return DB.Model(r)
}

// 预约状态
type ReservationStatus uint8

const (
	ReservationStatusPending ReservationStatus = iota + 1 // 进行中
	ReservationStatusSuccess                              // 已借阅
	ReservationStatusTimeout                              // 已超时
	ReservationStatusCancel                               // 已取消
)

// 预约状态文本
func PraseReservationStatus(status ReservationStatus) string {
	switch status {
	case ReservationStatusPending:
		return "进行中"
	case ReservationStatusSuccess:
		return "已借阅"
	case ReservationStatusTimeout:
		return "已超时"
	case ReservationStatusCancel:
		return "已取消"
	default:
		return "未知状态"
	}
}
