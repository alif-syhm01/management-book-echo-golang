package main

import (
	"management-book/configs"
	"management-book/models"
	"management-book/routes"
	"os"

	"github.com/labstack/echo/v4"
)

func init() {
	//configs.LoadEnv()
	configs.InitDatabase()
}

func main() {
	e := echo.New()

	e.Validator = models.NewCustomValidator(e)

	routes.InitRoute(e)

	//e.Logger.Fatal(e.Start(":8081"))
	e.Logger.Fatal(e.Start(GetPort()))
}

func GetPort() string {
	port := os.Getenv("PORT")
	return ":" + port
}
