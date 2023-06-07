package model

// 预约表（reservation）
// id (int，主键)
// book_id (int，外键，关联书籍表)
// reader_id (int，外键，关联读者表)

type Reservation struct {
	Model
	BookID     uint              `json:"book_id" gorm:"type:int;not null" binding:"required"`
	Book       Book              `json:"book" gorm:"foreignkey:BookID"`
	ReaderID   uint              `json:"reader_id" gorm:"type:int;not null" binding:"required"`
	Reader     Reader            `json:"reader" gorm:"foreignkey:ReaderID"`
	Status     ReservationStatus `json:"status" gorm:"type:tinyint;not null" binding:"required"`
	StatusText string            `json:"statusText" gorm:"-"` // 预约状态文本
}

// 预约状态
type ReservationStatus uint8

const (
	ReservationStatusPending ReservationStatus = iota + 1
	ReservationStatusSuccess
	ReservationStatusTimeout
	ReservationStatusCancel
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
