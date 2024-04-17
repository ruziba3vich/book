package service

import (
	"database/sql"
	"errors"

	"github.com/ruziba3vich/book/internal/app"
	"github.com/ruziba3vich/book/internal/authentication"
)

func Register(user app.User, db *sql.DB) (string, error) {
	query := "INSERT INTO users(id, username, email, password) VALUES ($1, $2, $3, $4);"
	var err error
	user.Password, err = authentication.HashPassword(user.Password)
	if err != nil {
		return "", errors.New("error while hashing the password")
	}
	_, error_ := db.Exec(query, user.ID, user.Username, user.Email, user.Password)
	if error_ != nil {
		return "", error_
	}
	token, e := authentication.GenerateJWTToken(user.ID)
	if e != nil {
		return "", e
	}
	return token, nil
}
