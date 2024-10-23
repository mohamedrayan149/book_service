package model

type Book struct {
	ID             string  `json:"id"`
	Title          string  `json:"title"`
	AuthorName     string  `json:"author_name"`
	Price          float64 `json:"price"`
	EbookAvailable bool    `json:"ebook_available"`
	PublishDate    string  `json:"publish_date"`
	Username       string  `json:"username"`
}

func NewBook() *Book {
	return &Book{
		ID:             "",
		Title:          "",
		AuthorName:     "",
		Price:          0.0,
		EbookAvailable: false,
		PublishDate:    "",
		Username:       "",
	}
}
