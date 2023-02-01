package model

import "time"

type User struct {
	Id        uint
	Name      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
