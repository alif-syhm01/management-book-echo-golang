package repositories

import (
	"management-book/configs"
	"management-book/models"
)

func GetBookAuthorGenres(bookGenreList *[]models.Book) error {
	// query find all books
	result := configs.DB.Preload("Author").Preload("Genre").Find(&bookGenreList)

	// check there result is getting error or not
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetBookAuthorGenre(bookGenre *models.Book, id string) error {
	// query find all books
	result := configs.DB.Preload("Author").Preload("Genre").First(&bookGenre, id)

	// check there result is getting error or not
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func AddBook(newBook *models.Book) error {
	// query to create book
	result := configs.DB.Create(&newBook)

	// check there result is getting error or not
	if result.Error != nil {
		return result.Error
	}

	// if save, just return nil
	return nil
}

func UpdateBook(updateBook *models.Book) error {
	// query to update book
	result := configs.DB.Save(&updateBook)

	// check there result is getting error or not
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteBook(deleteBook *models.Book, id string) error {
	// query to delete book
	result := configs.DB.Delete(&deleteBook, id)

	// check there result is getting error or not
	if result.Error != nil {
		return result.Error
	}

	// if save, just return nil
	return nil
}
