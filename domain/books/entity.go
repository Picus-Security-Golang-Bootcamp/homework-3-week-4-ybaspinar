package books

import (
	"gorm.io/gorm"
	"time"
)

type Books struct {
	gorm.Model
	Id          int64 `gorm:"primaryKey"`
	Pages       int64
	Stockamount int64
	Price       float64
	Stockid     int64
	Isbn        int64
	Author      Author `gorm:"embedded"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Author struct {
	gorm.Model
	Authorid   int64
	Authorname string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
