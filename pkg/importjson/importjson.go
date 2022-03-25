package importjson

import "encoding/json"

type Books []Book

func UnmarshalBooks(data []byte) (Books, error) {
	var r Books
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Books) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Book struct {
	BookID      uint    `json:"bookid"`
	Booktitle   string  `json:"booktitle"`
	Pages       int64   `json:"pages"`
	Stockamount int64   `json:"stockamount"`
	Price       float64 `json:"price"`
	Stockid     int64   `json:"stockid"`
	Isbn        int64   `json:"ısbn"`
	Author      Author  `json:"author"`
}

type Author struct {
	AuthorID   uint   `json:"authorid"`
	Authorname string `json:"authorname"`
}
