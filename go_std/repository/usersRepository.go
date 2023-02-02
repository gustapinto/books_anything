package repository

import (
	"database/sql"
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

func (r *UsersRepository) Find(userId uint, omitPassword bool) (user model.User, err error) {
	query := `SELECT * FROM "` + r.table + `" WHERE "id" = $1`
	if omitPassword {
		query = `SELECT "id", "name", "username", '' AS "password", "created_at", "updated_at" FROM "` + r.table + `" WHERE "id" = $1`
	}

	if err := r.db.QueryRow(query).Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return user, err
	}

	return user, nil
}

func (r *UsersRepository) GetAll(omitPassword bool) ([]model.User, error) {
	users := make([]model.User, 0)

	query := `SELECT * FROM "` + r.table + `"`
	if omitPassword {
		query = `SELECT "id", "name", "username", '' AS "password", "created_at", "updated_at" FROM "` + r.table + `"`
	}

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.User

		if err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (usersRepository *UsersRepository) PrepareUser(user *model.User) (model.User, error) {
	newUser := *user
	newUser.UpdatedAt = time.Now()

	if newUser.CreatedAt.IsZero() {
		newUser.CreatedAt = time.Now()
	}

	if newUser.Password != "" {
		hashedPassword, err := usersRepository.HashPassword(newUser.Password)
		if err != nil {
			return model.User{}, err
		}
		newUser.Password = hashedPassword
	}

	return newUser, nil
}

func (r *UsersRepository) Create(user *model.User) (model.User, error) {
	newUser, err := r.PrepareUser(user)
	if err != nil {
		return model.User{}, err
	}

	query := `
		INSERT INTO "` + r.table + `" ("name", "username", "password", "created_at", "updated_at")
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	if err := r.db.QueryRow(query, newUser.Name, newUser.Username, newUser.Password, newUser.CreatedAt, newUser.UpdatedAt).Scan(&newUser.Id); err != nil {
		return model.User{}, err
	}

	return newUser, nil
}

func (usersRepository *UsersRepository) Update(user *model.User) (newUser model.User, err error) {
	newUser, err = usersRepository.PrepareUser(user)
	if err != nil {
		return
	}

	query := `
		UPDATE "` + usersRepository.table + `"
		SET "name" = $1, "username" = $2, "password" = $3, "updated_at" = $4
		WHERE "id" = $5;
	`

	if _, err = usersRepository.db.Exec(query, newUser.Name, newUser.Username, newUser.Password, newUser.UpdatedAt, newUser.Id); err != nil {
		return
	}

	return newUser, nil
}

func (usersRepository *UsersRepository) Delete(userId uint) error {
	query := `
		DELETE FROM "` + usersRepository.table + `"
		WHERE "id" = $1;
	`

	if _, err := usersRepository.db.Exec(query, userId); err != nil {
		return err
	}

	return nil
}
