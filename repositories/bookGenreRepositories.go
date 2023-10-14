package repositories

import (
	"errors"
	"management-book/configs"
	"management-book/models"
)

func AddBookGenre(newBookGenre *models.Book) error {
	// return error when the newBookGenre is empty
	if len(newBookGenre.GenreId) == 0 {
		return errors.New("genreId is required")
	}

	// insert to table
	for _, genreId := range newBookGenre.GenreId {
		var bookGenre models.BookGenre

		bookGenre.BookId = newBookGenre.ID
		bookGenre.GenreId = uint(genreId)

		configs.DB.Create(&bookGenre)
	}

	return nil
}

func UpdateBookGenre(updateBookGenre *models.Book, id string) error {
	// only execute when the length of the genre id is more than zero
	if len(updateBookGenre.GenreId) > 0 {
		// find the genre id based on id book
		var bookGenres []models.BookGenre

		// find all book genre based on book id
		configs.DB.Where("book_id = ?", id).Find(&bookGenres)

		// force delete with data before then insert with new one
		configs.DB.Where("book_id = ?", id).Delete(&bookGenres)

		// insert new book genre
		for _, genreId := range updateBookGenre.GenreId {
			var newBookGenre models.BookGenre

			newBookGenre.BookId = updateBookGenre.ID
			newBookGenre.GenreId = uint(genreId)

			configs.DB.Create(&newBookGenre)
		}
	}

	return nil
}

func DeleteBookGenre(deleteBookGenre *models.Book, id string) error {

	// get all data
	var bookGenres []models.BookGenre

	// query to find all book_id by id book
	configs.DB.Where("book_id = ?", id).Find(&bookGenres)

	// loop the data by each one
	for _, bookGenre := range bookGenres {
		// delete the data
		result := configs.DB.Where("book_id = ?", id).Delete(&bookGenre)

		if result.Error != nil {
			return result.Error
		}
	}

	// if save, just return nil
	return nil
}
