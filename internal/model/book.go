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
	Title       string  `json:"title" gorm:"type:varchar(100);not null" binding:"required" example:"Go语言编程"`
	Author      string  `json:"author" gorm:"type:varchar(100);not null" binding:"required" example:"许式伟"`
	Publisher   string  `json:"publisher" gorm:"type:varchar(100);not null" binding:"required" example:"电子工业出版社"`
	PublishDate string  `json:"publishDate" gorm:"type:varchar(100);not null" binding:"required" example:"2019-01-01"`
	Summary     string  `json:"summary" gorm:"type:text;not null" binding:"required" example:"Go语言编程是一本介绍Go语言的书籍，内容包括Go语言的基础知识、并发编程、网络编程、Web编程、数据库编程等。"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2);not null" binding:"required" example:"99.99"`
	Stock       int     `json:"stock" gorm:"type:int;not null" binding:"required" example:"100"`
}
