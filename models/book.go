package models

type Book struct {
	ID              uint    `gorm:"primaryKey" json:"id" form:"id"`
	Title           string  `gorm:"not null" json:"title" form:"title" validate:"required"`
	PublicationYear int     `gorm:"not null" json:"publication_year" form:"publication_year" validate:"required,number"`
	Price           int     `gorm:"not null" json:"price" form:"price" validate:"required,number"`
	Author          Author  `json:"author"`
	AuthorID        uint    `json:"author_id" form:"author_id" validate:"required"`
	Genre           []Genre `gorm:"many2many:book_genres" json:"book_genres"`
	GenreId         []int   `gorm:"-" json:"genre_id" form:"genre_id"`
}

type BookResponse struct {
	ID              uint   `json:"id" form:"id"`
	Title           string `json:"title" form:"title"`
	PublicationYear int    `json:"publication_year" form:"publication_year"`
	Price           int    `json:"price" form:"price"`
}

type BookAuthorGenreResponse struct {
	ID              uint           `json:"id" form:"id"`
	Title           string         `json:"title" form:"title"`
	PublicationYear int            `json:"publication_year" form:"publication_year"`
	Price           int            `json:"price" form:"price"`
	Author          AuthorResponse `gorm:"-" json:"authors" form:"authors"`
	Genre           []Genre        `json:"book_genres" form:"book_genres"`
}

type BookCreateResponse struct {
	ID              uint   `json:"id" form:"id"`
	Title           string `json:"title" form:"title"`
	PublicationYear int    `json:"publication_year" form:"publication_year"`
	Price           int    `json:"price" form:"price"`
	AuthorID        uint   `json:"author_id" form:"author_id" `
	GenreId         []int  `json:"genre_id" form:"genre_id"`
}
