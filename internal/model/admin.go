package model

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 管理员表（admin）
// id (int，主键)
// name (varchar，管理员用户名)
// password (varchar，密码)

type Admin struct {
	Model
	Name     string `json:"name" gorm:"type:varchar(100);not null" binding:"required" example:"admin"`
	Password string `json:"password" gorm:"type:varchar(100);not null"`
}

func (m *Admin) Query() *gorm.DB {
	return DB.Model(m)
}

// EncryptPassword 加密密码
func (a *Admin) EncryptPassword() {
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("EncryptPassword error: %v", err)
	}
	a.Password = string(hash)
}

// ComparePassword 比较密码
func (a *Admin) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
	return err == nil
}
