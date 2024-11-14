package config

type BookUpdateData struct {
	ID    string `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
}
type BookAddData struct {
	Title          string  `json:"title" binding:"required"`
	AuthorName     string  `json:"author_name" binding:"required"`
	Price          float64 `json:"price" binding:"required"`
	EbookAvailable bool    `json:"ebook_available"`
	PublishDate    string  `json:"publish_date" binding:"required"`
}

type HttpError struct {
	Code int
}
