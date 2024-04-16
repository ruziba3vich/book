package app

import (
	"time"

	"github.com/google/uuid"
)

type Borrowing struct {
	ID         int       `json:"id"`
	UserId     uuid.UUID `json:"userId"`
	BookId     uuid.UUID `json:"bookId"`
	BorrowedAt time.Time `json:"borroweAt"`
	ReturnedAt time.Time `json:"returneAt"`
}
