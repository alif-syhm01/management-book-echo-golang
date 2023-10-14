package repositories

import (
	"management-book/configs"
	"management-book/models"
)

func GetGenres(genresList *[]models.Genre) error {
	// query find all genres
	result := configs.DB.Find(&genresList)

	// check there result is getting error or not
	if result.Error != nil {
		return result.Error
	}

	// if save, just return nil
	return nil
}

func GetGenre(genre *models.Genre, id string) error {
	// query find all genres
	result := configs.DB.First(&genre, id)

	// check there result is getting error or not
	if result.Error != nil {
		return result.Error
	}

	// if save, just return nil
	return nil
}

func AddGenre(newGenre *models.Genre) error {
	// query to create genre
	result := configs.DB.Create(&newGenre)

	// check there result is getting error or not
	if result.Error != nil {
		return result.Error
	}

	// is save, just return nil
	return nil
}

func UpdateGenre(updateGenre *models.Genre) error {
	// query to update genre
	result := configs.DB.Save(&updateGenre)

	// check there result is getting error or not
	if result.Error != nil {
		return result.Error
	}

	// is save, just return nil
	return nil
}

func DeleteGenre(deleteGenre *models.Genre, id string) error {
	// query to delete genre
	result := configs.DB.Delete(&deleteGenre, id)

	// check there result is getting error or not
	if result.Error != nil {
		return result.Error
	}

	// is save, just return nil
	return nil
}
