package model

import (
	"time"

	"github.com/google/uuid"
)

type AuthorInputModel struct {
	Name   string    `json:"name"`
	UserId uuid.UUID `json:"user_id"`
}

type AuthorViewModel struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	UserId    uuid.UUID `json:"user_id"`
}
