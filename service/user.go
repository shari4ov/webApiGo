package service

import (
	dto "home/DTO"
	"home/model"
	"home/storage"
	"log"

	"github.com/labstack/echo/v4"
)

func GetUserByID(id int) (model.User, error) {
	sqlStatement := `SELECT * FROM users WHERE id=$1;`
	db := storage.OpenConnection()
	row := db.QueryRow(sqlStatement, id)
	user := model.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Surname)
	if err != nil {
		return user, err
	}
	return user, nil
}
func GetRepoUsers() ([]model.User, error) {
	db := storage.OpenConnection()
	row, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	users := []model.User{}
	for row.Next() {
		var user model.User
		if err := row.Scan(&user.Id, &user.Name, &user.Surname); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	if err = row.Err(); err != nil {
		return users, err
	}
	return users, nil
}
func CreateUser(c echo.Context) error {
	user := dto.UserDTO{}
	if err := c.Bind(&user); err != nil {
		log.Fatal(err)
		return err
	}
	db := storage.OpenConnection()
	sqlStatement := `INSERT INTO users (id,name,surname) VALUES($1,$2,$3);`
	_, err := db.Exec(sqlStatement, 3, user.Name, user.Surname)
	if err != nil {
		panic(err)
	}
	return nil
}
