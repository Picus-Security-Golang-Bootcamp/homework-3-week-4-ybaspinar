package books

import (
	"errors"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-ybaspinar/pkg/importjson"
	"gorm.io/gorm"
)

type BooksRepository struct {
	db *gorm.DB
}

func NewBooksepository(db *gorm.DB) *BooksRepository {
	return &BooksRepository{db: db}
}

func (b *BooksRepository) Migrations() {
	b.db.AutoMigrate(&Books{})
}
func (b *BooksRepository) InsertData(booklist importjson.Books) {
	for _, book := range booklist {
		b.db.Where(Books{ID: uint(book.BookID)}).
			Attrs(Books{ID: uint(book.BookID), title: book.Booktitle, Pages: book.Pages, Stockamount: book.Stockamount, Price: book.Price, Stockid: book.Stockid}).FirstOrCreate(&book)
	}
}
func (b *BooksRepository) ListAll() []Books {
	var books []Books
	b.db.Find(&books)
	return books
}
func (b *BooksRepository) Search(key string) []Books {
	var books []Books
	b.db.Where("title LIKE ?", key).First(&books)
	return books
}
func (b *BooksRepository) GetById(key int) []Books {
	var books []Books
	b.db.Where("id = ?", key).First(&books)
	return books
}
func (b *BooksRepository) GetBooksWithAuthor(key string) []Books {
	var books []Books
	b.db.Table("books").Where("authorname LIKE ?", key).Find(&books)
	return books
}
func (b *BooksRepository) Buy(key string, amount int) error {
	var book Books
	b.db.Table("books").Where("authorname LIKE ?", key).Find(&book)
	if amount <= int(book.Stockamount) {
		b.db.Table("books").Where("id = ?", book.ID).Updates(map[string]interface{}{"stockamount": int(book.Stockamount) - amount})
	}
	return errors.New("Not Enough stock")
}
