package books

import (
	"gorm.io/gorm"
	"time"
)

type Books struct {
	BookID      uint `gorm:"primaryKey,unique"`
	Booktitle   string
	Pages       int64
	Stockamount int64
	Price       float64
	Stockid     int64
	Isbn        int64
	AuthorID    uint `gorm:"foreignKey:AuthorID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Author struct {
	AuthorID   uint `gorm:"primaryKey,unique"`
	Authorname string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
