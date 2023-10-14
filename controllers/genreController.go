package controllers

import (
	"management-book/configs"
	"management-book/models"
	"management-book/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetGenresController(c echo.Context) error {
	var genres []models.Genre

	err := repositories.GetGenres(&genres)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    []models.Genre{},
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menampilkan data",
		Status:  true,
		Data:    genres,
	})
}

func GetGenreController(c echo.Context) error {
	// get param
	id := c.Param("id")

	var genre models.Genre

	err := repositories.GetGenre(&genre, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    []models.Genre{},
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menampilkan data",
		Status:  true,
		Data:    genre,
	})
}

func CreateGenreController(c echo.Context) error {
	// struct models Genres
	var newGenre models.Genre

	// binding
	c.Bind(&newGenre)

	// error validate
	errValidate := c.Validate(newGenre)

	if errValidate != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: errValidate.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// call repositories add Genre
	err := repositories.AddGenre(&newGenre)

	// returning error if true
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// return success message
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menambahkan data",
		Status:  true,
		Data:    newGenre,
	})
}

func UpdateGenreController(c echo.Context) error {
	// get the parameter url
	id := c.Param("id")

	// make a variable for containing data from Genre model
	var updateGenre models.Genre

	// get data based on id
	configs.DB.First(&updateGenre, id)

	// binding the user input
	c.Bind(&updateGenre)

	// error validate
	errValidate := c.Validate(updateGenre)

	if errValidate != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: errValidate.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// call repositories update Genre
	err := repositories.UpdateGenre(&updateGenre)

	// check the error
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// return the success condition
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil mengubah data",
		Status:  true,
		Data:    updateGenre,
	})
}

func DeleteGenreController(c echo.Context) error {
	// get the parameter url
	id := c.Param("id")

	// make a variable for containing data from Genre model
	var Genre models.Genre

	// call repositories to delete data
	err := repositories.DeleteGenre(&Genre, id)

	// check the error
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// return the success condition
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menghapus data",
		Status:  true,
		Data:    nil,
	})
}
