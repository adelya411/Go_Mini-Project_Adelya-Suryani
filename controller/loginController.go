package controller

import (
	"net/http"
	"tulispuisi/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func LoginUser(c echo.Context) error {
	requestUser := new(models.User)
	if err := c.Bind(requestUser); err != nil {
		return err
	}

	db := c.Get("db").(*gorm.DB)
	user, err := models.FindUserByUsername(db, requestUser.Username)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err := user.CheckPassword(requestUser.Password); err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	return c.String(http.StatusOK, "Login berhasil")
}
