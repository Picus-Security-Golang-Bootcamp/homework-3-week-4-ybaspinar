package main

import (
	"fmt"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-ybaspinar/domain/books"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-ybaspinar/pkg/importjson"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-ybaspinar/pkg/newpsqldb"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	path, err := filepath.Abs("assets/books.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//Reads JSON file
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Books, err := importjson.UnmarshalBooks(file)
	println(Books[1].BookID)

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := newpsqldb.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init:", err)
	}
	log.Println("Postgres connected")
	booksRepo := books.NewBooksepository(db)
	booksRepo.Migrations()
	booksRepo.InsertData(Books)
}
