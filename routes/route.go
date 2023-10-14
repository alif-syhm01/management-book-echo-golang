package routes

import (
	"management-book/controllers"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.Use(middleware.Logger())
	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	// Genres Route
	eAuth.GET("/genres", controllers.GetGenresController)
	eAuth.GET("/genres/:id", controllers.GetGenreController)
	eAuth.POST("/genres", controllers.CreateGenreController)
	eAuth.PUT("/genres/:id", controllers.UpdateGenreController)
	eAuth.DELETE("/genres/:id", controllers.DeleteGenreController)

	// Authors Route
	eAuth.GET("/authors", controllers.GetAuthorsController)
	eAuth.GET("/authors/:id", controllers.GetAuthorController)
	eAuth.GET("/author-books", controllers.GetAuthorBooksController)
	eAuth.GET("/author-books/:id", controllers.GetAuthorBookController)
	eAuth.POST("/authors", controllers.CreateAuthorController)
	eAuth.PUT("/authors/:id", controllers.UpdateAuthorController)
	eAuth.DELETE("/authors/:id", controllers.DeleteAuthorController)

	// Books Route
	eAuth.GET("/books", controllers.GetBookAuthorGenresController)
	eAuth.GET("/books/:id", controllers.GetBookAuthorGenreController)
	eAuth.POST("/books", controllers.CreateBookController)
	eAuth.PUT("/books/:id", controllers.UpdateBookController)
	eAuth.DELETE("/books/:id", controllers.DeleteBookController)

	// Auth Route (because we want take the token, so we don't have to include the credentials in this request)
	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)
}
