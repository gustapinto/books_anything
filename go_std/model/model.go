package model

import "time"

type ModelInterface interface {
	Table() string
}

type Model struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
