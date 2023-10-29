package controller

import (
	"net/http"
	"tulispuisi/config"
	"tulispuisi/models"

	"github.com/labstack/echo/v4"
)

//Menambahkan kategori
func CreateKategori(c echo.Context) error {
	kategori := new(models.Kategori)
	if err := c.Bind(kategori); err != nil {
		return err
	}

	// Simpan kategori ke database
	if err := config.DB.Create(&kategori).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create kategori"})
	}

	return c.JSON(http.StatusCreated, kategori)
}

//Mendapatkan kategori
func GetKategori(c echo.Context) error {
    id := c.Param("id")
    kategori := new(models.Kategori)

    if err := config.DB.First(&kategori, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Kategori not found"})
    }

    return c.JSON(http.StatusOK, kategori)
}