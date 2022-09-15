package controller

import (
	"home/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users, _ := service.GetRepoUsers()
	return c.JSON(http.StatusOK, users)
}
func GetUserID(c echo.Context) error {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	user, err := service.GetUserByID(intId)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, user)
}
func CreateNewUser(c echo.Context) error {
	err := service.CreateUser(c)
	if err != nil {
		log.Fatal(err)
	}
	return c.String(http.StatusCreated, "create")

}
