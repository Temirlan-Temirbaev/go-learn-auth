package main

import (
	"github.com/labstack/echo/v4"
	"learn.com/config"
	"learn.com/controllers"
	"learn.com/models"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	e := echo.New()

	controllers.InitRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
