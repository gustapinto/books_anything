package repository

import (
	"database/sql"
	"time"

	"github.com/gustapinto/books_rest/go_std/model"
)

type BooksRepository struct {
	db                *sql.DB
	authorsRepository *AuthorsRepository
	usersRepository   *UsersRepository
}

func NewBooksRepository(db *sql.DB, authorsRepository *AuthorsRepository,
	usersRepository *UsersRepository) *BooksRepository {
	return &BooksRepository{
		db:                db,
		authorsRepository: authorsRepository,
		usersRepository:   usersRepository,
	}
}

func (_ *BooksRepository) Model() model.ModelInterface {
	return new(model.Book)
}

func (r *BooksRepository) Find(bookId uint) (book model.Book, err error) {
	query := `SELECT * FROM "` + r.Model().Table() + `" WHERE "id" = $1`

	if err := r.db.QueryRow(query, bookId).Scan(book.Dest()...); err != nil {
		return book, err
	}

	book.Author, err = r.authorsRepository.Find(book.Author.Id)
	if err != nil {
		return book, err
	}
	book.Creator, err = r.usersRepository.Find(book.Creator.Id)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *BooksRepository) All() ([]model.Book, error) {
	query := `SELECT * FROM "` + r.Model().Table() + `"`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	books := make([]model.Book, 0)

	for rows.Next() {
		var book model.Book

		if err := rows.Scan(book.Dest()...); err != nil {
			return nil, err
		}

		book.Author, err = r.authorsRepository.Find(book.Author.Id)
		if err != nil {
			return nil, err
		}
		book.Creator, err = r.usersRepository.Find(book.Creator.Id)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (r *BooksRepository) Create(book *model.Book) (newBook model.Book, err error) {
	query := `INSERT INTO "` + r.Model().Table() + `" ("isbn", "name", "author_id", "creator_id", "created_at", "updated_at")
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;`

	newBook = *book
	newBook.CreatedAt = time.Now()
	newBook.UpdatedAt = time.Now()

	if err := r.db.QueryRow(query, newBook.Fillable()...).Scan(&newBook.Id); err != nil {
		return newBook, err
	}

	newBook.Author, err = r.authorsRepository.Find(newBook.Author.Id)
	if err != nil {
		return newBook, err
	}
	newBook.Creator, err = r.usersRepository.FindWithoutPassword(newBook.Creator.Id)
	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (r *BooksRepository) Update(bookId uint, book *model.Book) (newBook model.Book, err error) {
	query := `UPDATE "` + r.Model().Table() + `"
		SET "isbn" = $1, "name" = $2, "author_id" = $3, "updated_at" = $4
		WHERE "id" = $5
		RETURNING "created_at"`

	newBook = *book
	newBook.Id = bookId
	newBook.UpdatedAt = time.Now()

	args := []any{
		newBook.ISBN,
		newBook.Name,
		newBook.Author.Id,
		newBook.UpdatedAt,
		newBook.Id,
	}

	if err = r.db.QueryRow(query, args...).Scan(&newBook.CreatedAt); err != nil {
		return
	}

	newBook.Author, err = r.authorsRepository.Find(newBook.Author.Id)
	if err != nil {
		return newBook, err
	}
	newBook.Creator, err = r.usersRepository.FindWithoutPassword(newBook.Creator.Id)
	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (r *BooksRepository) Delete(bookId uint) error {
	query := `DELETE FROM "` + r.Model().Table() + `" WHERE "id" = $1`

	if _, err := r.db.Exec(query, bookId); err != nil {
		return err
	}

	return nil
}
