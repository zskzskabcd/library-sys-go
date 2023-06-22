package model

import "gorm.io/gorm"

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
	Title       string  `json:"title" gorm:"type:varchar(100);not null;index" binding:"required" example:"Go语言编程"`                                          // 书名
	Author      string  `json:"author" gorm:"type:varchar(100);not null" binding:"" example:"许式伟"`                                                          // 作者名
	Publisher   string  `json:"publisher" gorm:"type:varchar(100);not null" binding:"" example:"电子工业出版社"`                                                   // 出版社
	PublishDate string  `json:"publishDate" gorm:"type:varchar(100);not null" binding:"" example:"2019-01-01"`                                              // 出版日期
	Summary     string  `json:"summary" gorm:"type:text;not null" binding:"" example:"Go语言编程是一本介绍Go语言的书籍，内容包括Go语言的基础知识、并发编程、网络编程、Web编程、数据库编程等。"`            // 书籍简介
	Cover       string  `json:"cover" gorm:"type:varchar(255);not null" binding:"" example:"https://img3.doubanio.com/view/subject/l/public/s29710665.jpg"` // 封面
	Price       float64 `json:"price" gorm:"type:decimal(10,2)" binding:"" example:"99.99"`                                                                 // 价格
	ISBN        string  `json:"isbn" gorm:"type:varchar(100);not null;index" binding:"" example:"9787121324947"`                                            // ISBN
	Stock       int     `json:"stock" gorm:"type:int" binding:"required" example:"100"`                                                                     // 库存量
	OriginStock int     `json:"originStock" gorm:"type:int" example:"100"`                                                                                  // 原始库存量
}

func (m *Book) Query() *gorm.DB {
	return DB.Model(m)
}
