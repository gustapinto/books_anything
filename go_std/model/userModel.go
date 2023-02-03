package model

import "fmt"

type User struct {
	Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

func (user *User) Table() string {
	return "users"
}

func (user *User) Migrate() string {
	query := `
		CREATE TABLE IF NOT EXISTS %[1]s (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255),
			username VARCHAR(100),
			password VARCHAR(255),
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		);

		ALTER TABLE "%[1]s" DROP CONSTRAINT IF EXISTS username_unique;

		ALTER TABLE "%[1]s" ADD CONSTRAINT username_unique UNIQUE (username);
	`

	return fmt.Sprintf(query, user.Table())
}
