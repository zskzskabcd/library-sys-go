package model

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 读者表（reader）
// id (int，主键)
// name (varchar，用户名)
// gender (varchar，性别)
// phone (varchar，手机号码)
type Reader struct {
	Model
	Name      string `json:"name" gorm:"type:varchar(100);not null" binding:"required" example:"张三"`           // 用户名
	Gender    string `json:"gender" gorm:"type:varchar(6);not null" binding:"required" example:"男"`            // 性别
	Phone     string `json:"phone" gorm:"type:varchar(20);not null" binding:"required" example:"18888888888"`  // 手机号码
	StudentNo uint   `json:"studentNo" gorm:"type:int;not null" binding:"required,min=0" example:"2018000000"` // 学号
	Key       string `json:"key" gorm:"type:varchar(100);not null" example:"123456"`                           // 密码（哈希值）
}

// Query 查询
func (r *Reader) Query() *gorm.DB {
	return DB.Model(r)
}

// EncryptPassword 加密密码
func (r *Reader) EncryptPassword() {
	hash, err := bcrypt.GenerateFromPassword([]byte(r.Key), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("EncryptPassword error: %v", err)
	}
	r.Key = string(hash)
}

// ComparePassword 比较密码
func (r *Reader) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(r.Key), []byte(password))
	return err == nil
}
