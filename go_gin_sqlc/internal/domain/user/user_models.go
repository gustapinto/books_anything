package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"password,omitempty"`
}

type UserPagination struct {
	Data        []User `json:"data"`
	TotalCount  uint   `json:"total_count"`
	TotalPages  uint   `json:"total_pages"`
	CurrentPage uint   `json:"current_page"`
	NextPage    uint   `json:"next_page"`
}
