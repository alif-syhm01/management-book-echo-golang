package controllers

import (
	"management-book/configs"
	"management-book/models"
	"management-book/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAuthorsController(c echo.Context) error {
	var authors []models.Author

	err := repositories.GetAuthors(&authors)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    []models.Author{},
		})
	}

	var authorResponse []models.AuthorResponse

	if len(authors) > 0 {
		// Initialize the authorResponse slice
		authorResponse = make([]models.AuthorResponse, len(authors))

		// Loop over authors and populate authorResponse
		for row, author := range authors {
			authorResponse[row].ID = author.ID
			authorResponse[row].Name = author.Name
			authorResponse[row].Nationality = author.Nationality
			authorResponse[row].Biography = author.Biography
		}
	} else {
		// insert empty slice of author
		authorResponse = []models.AuthorResponse{}
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menampilkan data",
		Status:  true,
		Data:    authorResponse,
	})
}

func GetAuthorController(c echo.Context) error {
	// get param
	id := c.Param("id")

	var author models.Author

	err := repositories.GetAuthor(&author, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    []models.Author{},
		})
	}

	var authorResponse models.AuthorResponse

	authorResponse.ID = author.ID
	authorResponse.Name = author.Name
	authorResponse.Nationality = author.Nationality
	authorResponse.Biography = author.Biography

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menampilkan data",
		Status:  true,
		Data:    authorResponse,
	})
}

func GetAuthorBooksController(c echo.Context) error {
	var authorBooks []models.Author

	err := repositories.GetAuthorBooks(&authorBooks)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    []models.Author{},
		})
	}

	var authorBookResponse []models.AuthorBookResponse

	if len(authorBooks) > 0 {
		// Initialize the authorBookResponse slice
		authorBookResponse = make([]models.AuthorBookResponse, len(authorBooks))

		// Loop over authorBook and populate authorBookResponse
		for row, authorBook := range authorBooks {
			// author response
			authorBookResponse[row].ID = authorBook.ID
			authorBookResponse[row].Name = authorBook.Name
			authorBookResponse[row].Nationality = authorBook.Nationality
			authorBookResponse[row].Biography = authorBook.Biography

			// book response
			var bookResponse []models.BookResponse

			if len(authorBook.Book) > 0 {
				// Initialize the bookResponse slice
				bookResponse = make([]models.BookResponse, len(authorBook.Book))

				// Loop over book and populate bookResponse
				for row, book := range authorBook.Book {
					bookResponse[row].ID = book.ID
					bookResponse[row].Title = book.Title
					bookResponse[row].PublicationYear = book.PublicationYear
					bookResponse[row].Price = book.Price
				}
			} else {
				bookResponse = []models.BookResponse{}
			}

			authorBookResponse[row].Book = bookResponse
		}
	} else {
		// insert empty slice of author
		authorBookResponse = []models.AuthorBookResponse{}
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menampilkan data",
		Status:  true,
		Data:    authorBookResponse,
	})
}

func GetAuthorBookController(c echo.Context) error {
	// get param
	id := c.Param("id")

	var authorBook models.Author

	err := repositories.GetAuthorBook(&authorBook, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    []models.Author{},
		})
	}

	// author response
	var authorBookResponse models.AuthorBookResponse

	authorBookResponse.ID = authorBook.ID
	authorBookResponse.Name = authorBook.Name
	authorBookResponse.Nationality = authorBook.Nationality
	authorBookResponse.Biography = authorBook.Biography

	// book response
	var bookResponse []models.BookResponse

	if len(authorBook.Book) > 0 {
		// Initialize the bookResponse slice
		bookResponse = make([]models.BookResponse, len(authorBook.Book))

		// Loop over book and populate bookResponse
		for row, book := range authorBook.Book {
			bookResponse[row].ID = book.ID
			bookResponse[row].Title = book.Title
			bookResponse[row].PublicationYear = book.PublicationYear
			bookResponse[row].Price = book.Price
		}
	} else {
		bookResponse = []models.BookResponse{}
	}

	authorBookResponse.Book = bookResponse

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menampilkan data",
		Status:  true,
		Data:    authorBookResponse,
	})
}

func CreateAuthorController(c echo.Context) error {
	// struct models Authors
	var newAuthor models.Author

	// binding
	c.Bind(&newAuthor)

	// error validate
	errValidate := c.Validate(newAuthor)

	if errValidate != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: errValidate.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// call repositories add Author
	err := repositories.AddAuthor(&newAuthor)

	// returning error if true
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	var authorResponse models.AuthorResponse

	authorResponse.ID = newAuthor.ID
	authorResponse.Name = newAuthor.Name
	authorResponse.Nationality = newAuthor.Nationality
	authorResponse.Biography = newAuthor.Biography

	// return success message
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menambahkan data",
		Status:  true,
		Data:    authorResponse,
	})
}

func UpdateAuthorController(c echo.Context) error {
	// get the parameter url
	id := c.Param("id")

	// make a variable for containing data from Author model
	var updateAuthor models.Author

	// get data based on id
	configs.DB.First(&updateAuthor, id)

	// binding the user input
	c.Bind(&updateAuthor)

	// error validate
	errValidate := c.Validate(updateAuthor)

	if errValidate != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: errValidate.Error(),
			Status:  false,
			Data:    nil,
		})
	}

	// call repositories update Author
	err := repositories.UpdateAuthor(&updateAuthor)

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
		Data:    updateAuthor,
	})
}

func DeleteAuthorController(c echo.Context) error {
	// get the parameter url
	id := c.Param("id")

	// make a variable for containing data from Author model
	var Author models.Author

	// call repositories to delete data
	err := repositories.DeleteAuthor(&Author, id)

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
