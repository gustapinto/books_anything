package model

import (
	"fmt"
)

type Author struct {
	Model
	Name    string `json:"name"`
	Creator User   `json:"creator"`
}

func (author *Author) Table() string {
	return "authors"
}

func (author *Author) Migrate() string {
	query := `
		CREATE TABLE IF NOT EXISTS %[1]s (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255),
			creator_id INTEGER REFERENCES %[2]s (id),
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		);
	`

	return fmt.Sprintf(query, author.Table(), new(User).Table())
}
