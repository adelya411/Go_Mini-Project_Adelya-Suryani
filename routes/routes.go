package routes

import (
	"github.com/labstack/echo/v4"
	"tulispuisi/controller"
)

func Init() *echo.Echo {
	e := echo.New()

	//Pemanggilan
	e.POST("/register", controller.RegisterUser)
	e.POST("/login", controller.LoginUser)

	e.POST("/puisi", controller.CreatePuisi)

	e.POST("/kategori", controller.CreateKategori)
	e.GET("/kategori/:id", controller.GetKategori)

	e.POST("/ratings", controller.CreateRating)
	e.GET("/ratings/:id", controller.GetRating)

	return e
}
