package app

import (
	"time"

	"github.com/google/uuid"
)

type BookBorrowRequest struct {
	UserId     uuid.UUID `json:"userId"`
	BookId     uuid.UUID `json:"bookId"`
	BorrowedAt time.Time `json:"borrowedAt"`
	ReturnedAt time.Time `json:"returnedAt"`
}
