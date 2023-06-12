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
	ID        uint      `gorm:"primarykey" json:"id" example:"1" format:"int64"`
	CreatedAt time.Time `json:"createdAt" example:"2021-01-01 00:00:00"`
	UpdatedAt time.Time `json:"updatedAt" example:"2021-01-01 00:00:00"`
}

func (m *Model) Query() *gorm.DB {
	return DB.Model(m)
}
