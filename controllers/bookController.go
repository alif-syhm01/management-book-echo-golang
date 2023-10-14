package controllers

import (
	"management-book/configs"
	"management-book/models"
	"management-book/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBookAuthorGenresController(c echo.Context) error {
	var bookAuthorGenre []models.Book

	err := repositories.GetBookAuthorGenres(&bookAuthorGenre)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    []models.Book{},
		})
	}

	var bookAuthorGenreResponse []models.BookAuthorGenreResponse

	if len(bookAuthorGenre) > 0 {
		// Initialize the bookAuthorGenreResponse slice
		bookAuthorGenreResponse = make([]models.BookAuthorGenreResponse, len(bookAuthorGenre))

		// Loop over bookAuthorGenre and populate bookAuthorGenreResponse
		for row, bookAuthorGenre := range bookAuthorGenre {
			bookAuthorGenreResponse[row].ID = bookAuthorGenre.ID
			bookAuthorGenreResponse[row].Title = bookAuthorGenre.Title
			bookAuthorGenreResponse[row].PublicationYear = bookAuthorGenre.PublicationYear
			bookAuthorGenreResponse[row].Price = bookAuthorGenre.Price
			bookAuthorGenreResponse[row].Author.ID = bookAuthorGenre.Author.ID
			bookAuthorGenreResponse[row].Author.Name = bookAuthorGenre.Author.Name
			bookAuthorGenreResponse[row].Author.Nationality = bookAuthorGenre.Author.Nationality
			bookAuthorGenreResponse[row].Author.Biography = bookAuthorGenre.Author.Biography
			bookAuthorGenreResponse[row].Genre = bookAuthorGenre.Genre
		}
	} else {
		// insert empty slice of author
		bookAuthorGenreResponse = []models.BookAuthorGenreResponse{}
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menampilkan data",
		Status:  true,
		Data:    bookAuthorGenreResponse,
	})
}

func GetBookAuthorGenreController(c echo.Context) error {
	// get the parameter id
	id := c.Param("id")

	var bookAuthorGenre models.Book

	err := repositories.GetBookAuthorGenre(&bookAuthorGenre, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    []models.Book{},
		})
	}

	var bookAuthorGenreResponse models.BookAuthorGenreResponse

	bookAuthorGenreResponse.ID = bookAuthorGenre.ID
	bookAuthorGenreResponse.Title = bookAuthorGenre.Title
	bookAuthorGenreResponse.PublicationYear = bookAuthorGenre.PublicationYear
	bookAuthorGenreResponse.Price = bookAuthorGenre.Price
	bookAuthorGenreResponse.Author.ID = bookAuthorGenre.Author.ID
	bookAuthorGenreResponse.Author.Name = bookAuthorGenre.Author.Name
	bookAuthorGenreResponse.Author.Nationality = bookAuthorGenre.Author.Nationality
	bookAuthorGenreResponse.Author.Biography = bookAuthorGenre.Author.Biography
	bookAuthorGenreResponse.Genre = bookAuthorGenre.Genre

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menampilkan data",
		Status:  true,
		Data:    bookAuthorGenreResponse,
	})
}

func CreateBookController(c echo.Context) error {
	// struct models book
	var newBook models.Book

	// binding
	c.Bind(&newBook)

	// error validate
	errValidate := c.Validate(newBook)

	if errValidate != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: errValidate.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// call repository add book
	errAddBook := repositories.AddBook(&newBook)

	if errAddBook != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: errAddBook.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// call repository bookGenre
	errBookGenre := repositories.AddBookGenre(&newBook)

	if errBookGenre != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: errBookGenre.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// make a response for successful create book
	var bookCreateResponse models.BookCreateResponse

	// append the model to the response
	bookCreateResponse.ID = newBook.ID
	bookCreateResponse.Title = newBook.Title
	bookCreateResponse.PublicationYear = newBook.PublicationYear
	bookCreateResponse.Price = newBook.Price
	bookCreateResponse.AuthorID = newBook.AuthorID
	bookCreateResponse.GenreId = newBook.GenreId

	// return response
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menambah data",
		Status:  true,
		Data:    bookCreateResponse,
	})

}

func UpdateBookController(c echo.Context) error {
	// get the parameter url
	id := c.Param("id")

	// make a variable for containing data from book model
	var updateBook models.Book

	// get data based on id
	configs.DB.First(&updateBook, id)

	// binding the user input
	c.Bind(&updateBook)

	// error validate
	errValidate := c.Validate(updateBook)

	if errValidate != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: errValidate.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// call repositories update book
	err := repositories.UpdateBook(&updateBook)

	// check the error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// call repository bookGenre
	errBookGenre := repositories.UpdateBookGenre(&updateBook, id)

	if errBookGenre != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: errBookGenre.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// make a response for successful update book
	var bookCreateResponse models.BookCreateResponse

	// append the model to the response
	bookCreateResponse.ID = updateBook.ID
	bookCreateResponse.Title = updateBook.Title
	bookCreateResponse.PublicationYear = updateBook.PublicationYear
	bookCreateResponse.Price = updateBook.Price
	bookCreateResponse.AuthorID = updateBook.AuthorID
	bookCreateResponse.GenreId = updateBook.GenreId

	// return the success condition
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil mengubah data",
		Status:  true,
		Data:    updateBook,
	})
}

func DeleteBookController(c echo.Context) error {
	// get the parameter url
	id := c.Param("id")

	// make a variable for containing data from book model
	var book models.Book

	// call repositories to delete pivot table
	errPivot := repositories.DeleteBookGenre(&book, id)

	// check the error
	if errPivot != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: errPivot.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// call repositories to delete data
	err := repositories.DeleteBook(&book, id)

	// check the error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
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
