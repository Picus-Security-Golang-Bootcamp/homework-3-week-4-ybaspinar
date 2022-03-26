package main

import (
	"fmt"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-ybaspinar/domain/books"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-ybaspinar/pkg/newpsqldb"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//Get filepath
	path, err := filepath.Abs("assets/books.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//Open file from file path
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//Bind file to struct
	Books, err := books.UnmarshalBooks(file)
	//Load .env
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//Open db connection
	db, err := newpsqldb.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init:", err)
	}
	log.Println("Postgres connected")

	booksRepo := books.NewBooksepository(db)
	//Create tables if not exists
	booksRepo.Migrations()
	//Insert data if not exists
	booksRepo.InsertData(Books)

	r := booksRepo.Search("Chess")
	for _, book := range r {
		println(book.Booktitle)
	}
}
