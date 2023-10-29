package controller

import (
	"net/http"
	"tulispuisi/config"
	"tulispuisi/models"

	"github.com/labstack/echo/v4"
)

//Menambahkan rating
func CreateRating(c echo.Context) error {
	rating := new(models.Rating)
	if err := c.Bind(rating); err != nil {
		return err
	}

	// Simpan rating ke database
	if err := config.DB.Create(&rating).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create rating"})
	}

	return c.JSON(http.StatusCreated, rating)
}

//Mendapatkan rating
func GetRating(c echo.Context) error {
    id := c.Param("id")
    rating := new(models.Rating)
    
    if err := config.DB.Preload("Penulis").First(&rating, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Rating not found"})
    }

    return c.JSON(http.StatusOK, rating)
}