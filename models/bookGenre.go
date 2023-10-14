package models

type BookGenre struct {
	BookId  uint `json:"book_id" form:"book_id"`
	GenreId uint `json:"genre_id" form:"genre_id"`
}
