package model

import (
	"time"

	"gorm.io/gorm"
)

// 借阅表（lending）
// id (int，主键)
// book_id (int，外键，关联书籍表)
// reader_id (int，外键，关联读者表)
// lend_time (date，借出时间)
// return_time (date，预期归还时间)

type Lending struct {
	Model
	BookID    uint          `json:"bookId" gorm:"type:int;not null" binding:"required" example:"1"` // 书籍ID
	Book      Book          `json:"book" gorm:"foreignkey:book_id"`
	ReaderID  uint          `json:"readerId" gorm:"type:int;not null" binding:"required" example:"1"` // 读者ID
	Reader    Reader        `json:"reader" gorm:"foreignkey:reader_id"`
	LendTime  time.Time     `json:"lendTime" gorm:"type:date;not null" binding:"required"`   // 借出时间
	ReturnTim time.Time     `json:"returnTime" gorm:"type:date;not null" binding:"required"` // 预期归还时间
	Status    LendingStatus `json:"status" gorm:"type:tinyint;not null" binding:"required" example:"1" enums:"1,2,3" enumdes:"1:借出,2:已归还,3:违约"`
}

func (l *Lending) Query() *gorm.DB {
	return DB.Model(l)
}

// 借阅状态
type LendingStatus uint8

const (
	LendingStatusLending   LendingStatus = iota + 1 // 借出
	LendingStatusReturned                           // 已归还
	LendingStatusViolation                          // 违约
)
