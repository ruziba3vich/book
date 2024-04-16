package app

import "github.com/google/uuid"

type Author struct {
	ID       uuid.UUID `json:"id"`
	Fullname string    `json:"fullname"`
}
