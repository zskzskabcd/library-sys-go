package model

import "time"

// 借阅表（lending）
// id (int，主键)
// book_id (int，外键，关联书籍表)
// reader_id (int，外键，关联读者表)
// lend_time (date，借出时间)
// return_time (date，预期归还时间)

type Lending struct {
	Model
	BookID    uint          `json:"book_id" gorm:"type:int;not null"`
	Book      Book          `json:"book" gorm:"foreignkey:BookID"`
	ReaderID  uint          `json:"reader_id" gorm:"type:int;not null"`
	Reader    Reader        `json:"reader" gorm:"foreignkey:ReaderID"`
	LendTime  time.Time     `json:"lend_time" gorm:"type:date;not null"`
	ReturnTim time.Time     `json:"return_time" gorm:"type:date;not null"`
	Status    LendingStatus `json:"status" gorm:"type:tinyint;not null" binding:"required"`
}

// 借阅状态
type LendingStatus uint8

const (
	LendingStatusLending LendingStatus = iota + 1
	// 已归还
	LendingStatusReturned
	// 违约
	LendingStatusViolation
)
