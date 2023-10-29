package controller

import (
	"net/http"
	"tulispuisi/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	//Hash password sebelum menyimpan
	if err := user.HashPassword(); err != nil {
		return err
	}

	//Simpan User
	if err := models.CreateUser(c.Get("db").(*gorm.DB), user); err != nil {
		return err
	}

	//Hapus password dari respons
	user.Password = ""
	
	return c.JSON(http.StatusCreated, user)
}
