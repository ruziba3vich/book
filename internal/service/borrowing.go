package service

import (
	"database/sql"
	"errors"
	"log"
	"time"

	m "github.com/ruziba3vich/book/internal/app"
)

func BorrowBook(bookBorrowRequest m.BookBorrowRequest, db *sql.DB) (responseObj *m.Borrowing, err error) {
	bookBorrowRequest.BorrowedAt = time.Now()
	query := "INSERT INTO borrowed_books(book_id, user_id, borrowed_at) values ($1, $2, $3) RETURNING book_id, user_id, borrowed_at;"
	row := db.QueryRow(query, bookBorrowRequest.BookId, bookBorrowRequest.UserId, bookBorrowRequest.BorrowedAt)

	if err := row.Scan(&responseObj.BookId, &responseObj.UserId, &responseObj.BorrowedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Fatal("returned an empty row,", err)
		}
		log.Fatal(err)
	}
	return responseObj, nil
}

func ReturnBook(bookBorrowRequest m.BookBorrowRequest, db *sql.DB) error {
	bookBorrowRequest.ReturnedAt = time.Now()

	query := `
        UPDATE borrowed_books
        SET returned_at = CURRENT_TIMESTAMP
        WHERE book_id = $1 AND user_id = $2 AND returned_at IS NULL;
    `
	_, err := db.Exec(query, bookBorrowRequest.BookId, bookBorrowRequest.UserId)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
