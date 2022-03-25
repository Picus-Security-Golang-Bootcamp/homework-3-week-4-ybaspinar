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
	b.db.AutoMigrate(&Author{})

}
func (b *BooksRepository) InsertData(booklist importjson.Books, Authorlist importjson.Authors) {
	for _, author := range Authorlist {
		b.db.FirstOrCreate(&author, Author{AuthorID: author.AuthorID})
	}
	for _, book := range booklist {
		b.db.FirstOrCreate(&book, Books{Isbn: book.Isbn})
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
func (b *BooksRepository) GetAuthorWithBooks(key string) []Books {
	var books []Books
	b.db.Table("books").Where("authorname LIKE ?", key).Find(&books)
	return books
}
func (b *BooksRepository) Buy(key string, amount int) error {
	var book Books
	b.db.Table("books").Where("authorname LIKE ?", key).Find(&book)
	if amount <= int(book.Stockamount) {
		b.db.Table("books").Where("id = ?", book.BookID).Updates(map[string]interface{}{"stockamount": int(book.Stockamount) - amount})
	}
	return errors.New("Not Enough stock")
}
func (b *BooksRepository) Delete(key int) {
	b.db.Table("books").Delete(map[string]interface{}{"ID": key})
}
func (b *BooksRepository) FindDeleted(key string) []Books {
	var books []Books
	b.db.Unscoped().Where("title LIKE ?", key).Find(&books)
	return books
}
