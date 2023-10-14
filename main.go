package main

import (
	"management-book/configs"
	"management-book/models"
	"management-book/routes"

	"github.com/labstack/echo/v4"
)

func init() {
	configs.LoadEnv()
	configs.InitDatabase()
}

func main() {
	e := echo.New()

	e.Validator = models.NewCustomValidator(e)

	routes.InitRoute(e)

	e.Logger.Fatal(e.Start(":8081"))
}
