package model

import "fmt"

type Book struct {
	Model
	ISBN    string `json:"isbn"`
	Name    string `json:"name"`
	Author  Author `json:"author"`
	Creator User   `json:"creator"`
}

func (book *Book) Table() string {
	return "books"
}

func (book *Book) Migrate() string {
	query := `
		CREATE TABLE IF NOT EXISTS %[1]s (
			id SERIAL PRIMARY KEY,
			isbn CHAR(13) UNIQUE NOT NULL,
			name VARCHAR(255) UNIQUE NOT NULL,
			author_id INTEGER REFERENCES %[2]s (id) NOT NULL,
			creator_id INTEGER REFERENCES %[3]s (id) NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		);
	`

	return fmt.Sprintf(query, book.Table(), new(Author).Table(), new(User).Table())
}

func (book *Book) Dest() []any {
	return []any{
		&book.Id,
		&book.ISBN,
		&book.Name,
		&book.Author.Id,
		&book.Creator.Id,
		&book.CreatedAt,
		&book.UpdatedAt,
	}
}

func (book *Book) Fillable() []any {
	return []any{
		book.ISBN,
		book.Name,
		book.Author.Id,
		book.Creator.Id,
		book.CreatedAt,
		book.UpdatedAt,
	}
}
