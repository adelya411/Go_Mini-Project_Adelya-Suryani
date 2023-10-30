package main

import (
	// "time"
	"tulispuisi/config"
	"tulispuisi/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDataBase()

	e := echo.New()

	// Mengatur rute-rute API
	e.GET("/puisi", controller.GetPuisi)
	e.GET("/puisi/:id", controller.GetPuisiByID)
	e.POST("/puisi", controller.CreatePuisi)
	// e.PUT("/puisi/:id", controller.UpdatePuisi)
	e.DELETE("/puisi/:id", controller.DeletePuisi)

	e.Start(":8080")


}