package model

type Book struct {
	BookID string `json:"book_id"`
	Name   string `json:"name"`
}

func (Book) TableName() string {
	return "book"
}
