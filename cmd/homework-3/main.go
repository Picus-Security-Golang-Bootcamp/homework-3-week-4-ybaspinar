package main

import (
	"fmt"
	"github.com/Picus-Security-Golang-Bootcamp/homework-3-week-4-ybaspinar/pkg/importjson"
	"io/ioutil"
	"os"
)

func main() {
	var Books []importjson.Book
	//Reads JSON file
	file, err := ioutil.ReadFile("assets/books.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Books, err = importjson.UnmarshalBooks(file)

}
