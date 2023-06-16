package model

import (
	// "github.com/glebarez/sqlite"
	"library-sys-go/pkg/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Setup initializes the database instance
func Setup() {
	mdbConf := config.GetMysqlConfig()
	// db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(mdbConf.GetMysqlDSN()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Migration(db)
	DB = db
}

func migrate(db *gorm.DB, model interface{}) {
	err := db.AutoMigrate(model)
	if err != nil {
		panic("failed to migrate database + " + err.Error())
	}
}

// Migration migrate the schema
func Migration(db *gorm.DB) {
	migrate(db, &Book{})
	migrate(db, &Reservation{})
	migrate(db, &Reader{})
	migrate(db, &Admin{})
	migrate(db, &Lending{})
}

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id" example:"1" format:"int64"`                                     // 主键ID
	CreatedAt time.Time      `json:"createdAt" example:"2023-06-13T19:06:22.514+08:00"`                                   // 创建时间
	UpdatedAt time.Time      `json:"updatedAt" example:"2023-06-13T19:06:22.514+08:00"`                                   // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt" swaggertype:"string" example:"2023-06-13T19:06:22.514+08:00"` // 删除时间 - 软删除
}

func (m *Model) Query() *gorm.DB {
	return DB.Model(m)
}

func UnscopedQuery() *gorm.DB {
	return DB.Unscoped()
}
