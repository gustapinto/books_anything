package model

import "time"

type User struct {
	Id        uint      `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password,omitempty" db:"password"`
}
