package service

import (
	"database/sql"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/ruziba3vich/book/internal/app"
)

func CreateAuthor(auth app.Author, db *sql.DB) (responseObj *app.Author, err error) {
	auth.ID = uuid.New()
	query := "INSERT INTO authors(id, fullname) VALUES ($1, $2) RETURNING id, fullname;"
	row := db.QueryRow(query, auth.ID, auth.Fullname)
	if err := row.Scan(&responseObj.ID, &responseObj.Fullname); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Fatal("returning an empty line :", err)
		} else {
			log.Fatal(err)
		}
	}
	return responseObj, nil
}
