package model

// 管理员表（admin）
// id (int，主键)
// name (varchar，管理员用户名)
// password (varchar，密码)

type Admin struct {
	Model
	Name     string `json:"name" gorm:"type:varchar(100);not null" binding:"required" example:"admin"`
	Password string `json:"password" gorm:"type:varchar(100);not null"`
}
