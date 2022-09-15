package controller

import (
	"home/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users, _ := service.GetRepoUsers()
	return c.JSON(http.StatusOK, users)
}
