package repository

import (
	"errors"
	"time"

	"github.com/gustapinto/books_rest/go_gin_sqlx/pkg/auth"
	"github.com/gustapinto/books_rest/go_gin_sqlx/pkg/model"
	"github.com/jmoiron/sqlx"
)

var ErrInvalidAuthentication = errors.New("invalid authentication")

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Find(userId uint) (*model.User, error) {
	var user model.User

	query := `SELECT * FROM "users" WHERE id = $1;`
	if err := r.db.QueryRowx(query, userId).StructScan(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) All() ([]model.User, error) {
	var users []model.User

	if err := r.db.Select(&users, `SELECT * FROM "users"`); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Create(user model.User) (*model.User, error) {
	var newUser model.User

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	query := `INSERT INTO "users" (created_at, updated_at, name, username, password)
		VALUES (:created_at, :updated_at, :name, :username, :password)
		RETURNING *;`

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.StructScan(&newUser)
	}

	return &newUser, nil
}

func (r *UserRepository) Update(userId uint, user model.User) (*model.User, error) {
	var newUser model.User

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Id = userId
	user.Password = hashedPassword
	user.UpdatedAt = time.Now()

	query := `UPDATE "users"
		SET name = :name, username = :username, password = :password, updated_at = :updated_at
		WHERE id = :id
		RETURNING *`

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.StructScan(&newUser)
	}

	return &newUser, nil
}

func (r *UserRepository) Delete(userId uint) error {
	query := `DELETE FROM "users" WHERE id = $1`

	if _, err := r.db.Exec(query, userId); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Login(username, password string) (*model.User, error) {
	var user model.User

	query := `SELECT * FROM "users" WHERE "username" = $1`
	if err := r.db.QueryRowx(query, username).StructScan(&user); err != nil {
		return nil, err
	}

	if matched := auth.ComparePasswords(password, user.Password); !matched {
		return nil, ErrInvalidAuthentication
	}

	return &user, nil
}
