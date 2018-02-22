package model

type Book struct {
	BookID string `json:"book_id" gorm:"column:book_id; primary_key"`
	Name   string `json:"name" gorm:"column:name;"`
}

func (Book) TableName() string {
	return "book"
}
