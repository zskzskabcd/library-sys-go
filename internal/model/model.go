package model

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Setup initializes the database instance
func Setup() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Migration(db)
	DB = db
}

func migrate(db *gorm.DB, models ...interface{}) {
	err := db.AutoMigrate(models...)
	if err != nil {
		panic("failed to migrate database + " + err.Error())
	}
}

// Migration migrate the schema
func Migration(db *gorm.DB) {
	// Migrate the schema
	// migrate(db, &Book{}, &Reader{}, &Reservation{}, &Lending{})
	// migrate(db, &Admin{})
}

type Model struct {
	ID        uint   `gorm:"primarykey" json:"id" example:"1" format:"int64"`
	CreatedAt string `json:"createdAt" example:"2021-01-01 00:00:00"`
	UpdatedAt string `json:"updatedAt" example:"2021-01-01 00:00:00"`
}

func (m *Model) Query() *gorm.DB {
	return DB.Model(m)
}
