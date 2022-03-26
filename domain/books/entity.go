package books

import (
	"gorm.io/gorm"
	"time"
)

import "encoding/json"

type Books []Book
type Authors []Author

func UnmarshalBooks(data []byte) (Books, error) {
	var r Books
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Book) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Book struct {
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
	Author      Author         `gorm:"-"`
}

type Author struct {
	AuthorID   uint `gorm:"primaryKey,unique"`
	Authorname string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
