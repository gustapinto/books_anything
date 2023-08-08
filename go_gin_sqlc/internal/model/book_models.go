package model

import (
	"os/user"
	"time"

	"github.com/google/uuid"
)

type BookInputModel struct {
	ISBN   string    `json:"isbn"`
	Name   string    `json:"name"`
	UserId uuid.UUID `json:"user_id"`
}

type Book struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ISBN      string    `json:"isbn"`
	Name      string    `json:"name"`
	User      user.User `json:"user"`
	Author    Author    `json:"author"`
}
