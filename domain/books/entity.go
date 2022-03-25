package books

import (
	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	title       string
	Pages       int64
	Stockamount int64
	Price       float64
	Stockid     int64
	Isbn        int64
	Author      Author `gorm:"embedded"`
}

type Author struct {
	gorm.Model
	AuthorID   uint
	Authorname string
}
