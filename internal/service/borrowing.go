package service

import (
	"database/sql"
	"errors"
	"log"

	m "github.com/ruziba3vich/book/internal/app"
)

func BorrowBook(bookBorrowRequest m.BookBorrowRequest, db *sql.DB) (responseObj *m.Borrowing, err error) {

	query := "INSERT INTO borrowed_books(book_id, user_id, borrowed_at, returned_at) values ($1, $2, $3, $4) RETURNING book_id, user_id, borrowed_at, returned_at;"
	row := db.QueryRow(query, responseObj.BookId, responseObj.UserId, responseObj.BorrowedAt, responseObj.ReturnedAt)

	if err := row.Scan(&responseObj.BookId, &responseObj.UserId, &responseObj.BorrowedAt, responseObj.ReturnedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Fatal("returned an empty row,", err)
		}
		log.Fatal(err)
	}
	return responseObj, nil
}
