package controller

import (
	"net/http"
	"tulispuisi/config"
	"tulispuisi/models"

	"github.com/labstack/echo/v4"
)

func GetPuisi(c echo.Context) error {
	var puisi []models.Puisi
	config.DB.Find(&puisi)
	return c.JSON(http.StatusOK, puisi)
}

func GetPuisiByID(c echo.Context) error {
	id := c.Param("id")
	var puisi models.Puisi
	if err := config.DB.Where("id = ?", id).First(&puisi).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, puisi)
}

func CreatePuisi(c echo.Context) error {
	puisi := new(models.Puisi)
	if err := c.Bind(puisi); err != nil {
		return err
	}
	config.DB.Create(puisi)
	return c.JSON(http.StatusCreated, puisi)
}

// func UpdatePuisi(c echo.Context) error {
// 	id := c.Param("id")
// 	puisi := new(models.Puisi)
// 	if err := c.Bind(puisi); err != nil {
// 		return err
// 	}
// 	var existingPuisi models.Puisi
// 	if err := config.DB.Where("id = ?", id).First(&existingPuisi).Error; err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}
// 	config.DB.Update()
// 	return c.JSON(http.StatusOK, existingPuisi)
// }

func DeletePuisi(c echo.Context) error {
	id := c.Param("id")
	var puisi models.Puisi
	if err := config.DB.Where("id = ?", id).First(&puisi).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	config.DB.Delete(&puisi)
	return c.NoContent(http.StatusNoContent)
}
