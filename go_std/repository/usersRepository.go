package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gustapinto/books_rest/go_std/model"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepository struct {
	table string
	db    *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{
		table: "users",
		db:    db,
	}
}

func (r *UsersRepository) Migrate() string {
	return `
		CREATE TABLE IF NOT EXISTS ` + r.table + ` (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255),
			username VARCHAR(100),
			password VARCHAR(255),
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		)
	`
}

func (r *UsersRepository) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (r *UsersRepository) Create(user *model.User) (*model.User, error) {
	newUser := *user
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	hashedPassword, err := r.HashPassword(newUser.Password)
	if err != nil {
		return nil, err
	}

	newUser.Password = hashedPassword

	query := `
		INSERT INTO "` + r.table + `" ("name", "username", "password", "created_at", "updated_at")
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	fmt.Printf("%s", newUser.Password)

	if err := r.db.QueryRow(query, newUser.Name, newUser.Username, newUser.Password, newUser.CreatedAt, newUser.UpdatedAt).Scan(&newUser.Id); err != nil {
		return nil, err
	}

	return &newUser, nil
}
