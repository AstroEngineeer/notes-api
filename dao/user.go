package dao

import (
	"github.com/vigneshganesan008/notes-api/models"
)

func InsertUser(user models.User) (id int64, err error) {
	if err = db.QueryRow(
		"INSERT INTO users(username, password) VALUES($1, $2) RETURNING id;",
		user.Username,
		user.Password,
	).Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func GetUser(username string) (user models.User, err error) {
	if err := db.QueryRow(
		"SELECT id, username, password FROM users WHERE username=$1",
		username,
	).Scan(&user.ID, &user.Username, &user.Password); err != nil {
		return user, err
	}

	return user, nil
}
