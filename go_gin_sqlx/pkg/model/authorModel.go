package model

import "time"

type Author struct {
	Id        uint      `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	CreatedBy uint      `json:"created_by,omitempty" db:"created_by"`
}

type AuthorCreator struct {
	Author `db:"author"` // Omit json key to embedded data into json body
	User   `json:"creator" db:"user"`
}
