package model

import "time"

type Book struct {
	Id        uint      `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	ISBN      string    `json:"isbn" db:"isbn"`
	Name      string    `json:"name" db:"name"`
	AuthorId  uint      `json:"author_id" db:"author_id"`
	CreatedBy uint      `json:"created_by" db:"created_by"`
}

type BookAuthorCreator struct {
	Book   `db:"book"`
	Author `json:"author" db:"author"`
	User   `json:"creator" db:"user"`
}
