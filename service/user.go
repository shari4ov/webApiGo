package service

import (
	"home/model"
	"home/storage"
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
