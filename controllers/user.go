package controllers

import (
	"github.com/labstack/echo/v4"
	"learn.com/middleware"
	"learn.com/services"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/users", services.CreateUser)
	e.POST("/users/login", services.LoginUser)
	e.GET("/users", services.GetUsers)
	e.GET("/users/profile", services.GetUserData, middleware.JWTMiddleware)
	e.GET("/users/:id", services.GetUser)
}
