package repositories

import (
	"management-book/configs"
	"management-book/models"
)

func GetAuthors(authors *[]models.Author) error {
	// query find all author
	result := configs.DB.Find(&authors)

	// check there result is getting error or not
	if result.Error != nil {
		return result.Error
	}

	// if save, just return nil
	return nil
}

func GetAuthor(authors *models.Author, id string) error {
	// query find all author
	result := configs.DB.First(&authors, id)

	// check there result is getting error or not
	if result.Error != nil {
		return result.Error
	}

	// if save, just return nil
	return nil
}

func GetAuthorBooks(authorBooksList *[]models.Author) error {
	// query find all author book
	result := configs.DB.Preload("Book").Find(&authorBooksList)

	// check there result is getting error or not
	if result.Error != nil {
		return result.Error
	}

	// if save, just return nil
	return nil
}

func GetAuthorBook(authorBooksList *models.Author, id string) error {
	// query find all author book
	result := configs.DB.Preload("Book").First(&authorBooksList, id)

	// check there result is getting error or not
	if result.Error != nil {
		return result.Error
	}

	// if save, just return nil
	return nil
}

func AddAuthor(newAuthor *models.Author) error {
	// query to create author
	result := configs.DB.Create(&newAuthor)

	// check there result is getting error or not
	if result.Error != nil {
		return result.Error
	}

	// is save, just return nil
	return nil
}

func UpdateAuthor(updateAuthor *models.Author) error {
	// query to update author
	result := configs.DB.Save(&updateAuthor)

	// check there result is getting error or not
	if result.Error != nil {
		return result.Error
	}

	// is save, just return nil
	return nil
}

func DeleteAuthor(deleteAuthor *models.Author, id string) error {
	// query to delete author
	result := configs.DB.Delete(&deleteAuthor, id)

	// check there result is getting error or not
	if result.Error != nil {
		return result.Error
	}

	// is save, just return nil
	return nil
}
