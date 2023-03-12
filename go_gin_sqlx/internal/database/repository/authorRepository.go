package repository

import (
	"time"

	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/model"
	"github.com/jmoiron/sqlx"
)

const (
	queryAuthorWithCreator = `
		SELECT
			a.id AS "author.id",
			a.created_at AS "author.created_at",
			a.updated_at AS "author.updated_at",
			a."name" AS "author.name",
			u.id AS "user.id",
			u.created_at AS "user.created_at",
			u.updated_at AS "user.updated_at",
			u."name" AS "user.name",
			u.username AS "user.username"
		FROM
			authors a
		INNER JOIN users u ON
			u.id = a.created_by
	`
)

type AuthorRepository struct {
	db *sqlx.DB
}

func NewAuthorRepository(db *sqlx.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (r *AuthorRepository) Find(authorId uint) (*model.AuthorCreator, error) {
	var author model.AuthorCreator

	query := queryAuthorWithCreator + `WHERE a.id = $1`
	if err := r.db.Get(&author, query, authorId); err != nil {
		return nil, err
	}

	return &author, nil
}

func (r *AuthorRepository) All() ([]model.AuthorCreator, error) {
	var authors []model.AuthorCreator

	if err := r.db.Select(&authors, queryAuthorWithCreator); err != nil {
		return nil, err
	}

	return authors, nil
}

func (r *AuthorRepository) Create(author model.Author) (*model.AuthorCreator, error) {
	var newAuthor model.AuthorCreator

	author.CreatedAt = time.Now()
	author.UpdatedAt = time.Now()

	query := `
		WITH authors AS (
			INSERT INTO authors
				(created_at, updated_at, name, created_by)
			VALUES
				(:created_at, :updated_at, :name, :created_by)
			RETURNING *
		)
	` + queryAuthorWithCreator
	rows, err := r.db.NamedQuery(query, author)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.StructScan(&newAuthor); err != nil {
			return nil, err
		}
	}

	return &newAuthor, err
}

func (r *AuthorRepository) Update(authorId uint, author model.Author) (*model.AuthorCreator, error) {
	var newAuthor model.AuthorCreator

	author.Id = authorId
	author.UpdatedAt = time.Now()

	query := `
		WITH authors AS (
			UPDATE authors
			SET
				name = :name
			WHERE
				id = :id
			RETURNING *
		)
	` + queryAuthorWithCreator
	rows, err := r.db.NamedQuery(query, author)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.StructScan(&newAuthor); err != nil {
			return nil, err
		}
	}

	return &newAuthor, nil
}

func (r *AuthorRepository) Delete(authorId uint) error {
	query := `DELETE FROM authors WHERE id = $1`
	if _, err := r.db.Exec(query, authorId); err != nil {
		return err
	}

	return nil
}
