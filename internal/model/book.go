package model

// 书籍表（book）
// id (int, 主键)
// title (varchar，书名)
// author (varchar，作者名)
// publisher (varchar，出版社)
// publish_date (date，出版日期)
// summary (text，书籍简介)
// price (decimal，价格)
// stock (int，库存量)

type Book struct {
	Model
	Title       string  `json:"title" gorm:"type:varchar(100);not null" binding:"required"`
	Author      string  `json:"author" gorm:"type:varchar(100);not null" binding:"required"`
	Publisher   string  `json:"publisher" gorm:"type:varchar(100);not null" binding:"required"`
	PublishDate string  `json:"publishDate" gorm:"type:varchar(100);not null" binding:"required"`
	Summary     string  `json:"summary" gorm:"type:text;not null" binding:"required"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2);not null" binding:"required"`
	Stock       int     `json:"stock" gorm:"type:int;not null" binding:"required"`
}
