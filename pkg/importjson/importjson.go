package importjson

import "encoding/json"

type Books []Book
type Authors []Author

func UnmarshalBooks(data []byte) (Books, Authors, error) {
	var r Books
	var a Authors
	err := json.Unmarshal(data, &r)
	err = json.Unmarshal(data, &a)
	return r, a, err
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
	AuthorID    uint    `gorm:"foreignKey:AuthorID"`
}

type Author struct {
	AuthorID   uint   `json:"authorid"`
	Authorname string `json:"authorname"`
}
