package repository

import "database/sql"

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
