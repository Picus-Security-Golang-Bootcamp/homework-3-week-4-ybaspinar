package books

import "gorm.io/gorm"

type BooksRepository struct {
	db *gorm.DB
}

func NewBooksepository(db *gorm.DB) *BooksRepository {
	return &BooksRepository{db: db}
}

func (b *BooksRepository) Migrations() {
	b.db.AutoMigrate(&Books{})
}
