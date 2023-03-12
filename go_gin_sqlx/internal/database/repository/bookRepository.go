package repository

import (
	"time"

	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/model"
	"github.com/jmoiron/sqlx"
)

const queryBookWithAuthorAndCreator = `
	SELECT
		b.id AS "book.id",
		b.created_at AS "book.created_at",
		b.updated_at AS "book.updated_at",
		b.isbn AS "book.isbn",
		b."name" AS "book.name",
		b.author_id AS "book.author_id",
		b.created_by AS "book.created_by",
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
		books b
	INNER JOIN authors a ON
		a.id = b.author_id
	INNER JOIN users u ON
		u.id = b.created_by
`

type BookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) Find(bookId uint) (*model.BookAuthorCreator, error) {
	var book model.BookAuthorCreator

	query := queryBookWithAuthorAndCreator + `WHERE b.id = $1`
	if err := r.db.Get(&book, query, bookId); err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *BookRepository) All() ([]model.BookAuthorCreator, error) {
	var books []model.BookAuthorCreator

	if err := r.db.Select(&books, queryBookWithAuthorAndCreator); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepository) Create(book model.Book) (*model.BookAuthorCreator, error) {
	var newBook model.BookAuthorCreator

	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	query := `
		WITH books AS (
			INSERT INTO books
				(created_at, updated_at, isbn, name, author_id, created_by)
			VALUES
				(:created_at, :updated_at, :isbn, :name, :author_id, :created_by)
			RETURNING *
		)
	` + queryBookWithAuthorAndCreator
	rows, err := r.db.NamedQuery(query, book)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.StructScan(&newBook); err != nil {
			return nil, err
		}
	}

	return &newBook, err
}

func (r *BookRepository) Update(bookId uint, book model.Book) (*model.BookAuthorCreator, error) {
	var newBook model.BookAuthorCreator

	book.Id = bookId
	book.UpdatedAt = time.Now()

	query := `
		WITH books AS (
			UPDATE books
			SET
				isbn = :isbn,
				name = :name,
				author_id = :author_id
			WHERE
				id = :id
			RETURNING *
		)
	` + queryBookWithAuthorAndCreator
	rows, err := r.db.NamedQuery(query, book)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.StructScan(&newBook); err != nil {
			return nil, err
		}
	}

	return &newBook, nil

}

func (r *BookRepository) Delete(bookId uint) error {
	query := `DELETE FROM "books" WHERE id = $1`
	if _, err := r.db.Exec(query, bookId); err != nil {
		return err
	}

	return nil
}
