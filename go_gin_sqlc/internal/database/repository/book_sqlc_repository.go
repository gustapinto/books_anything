package repository

import (
	"github.com/google/uuid"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/model"
)

type BookSqlcRepository struct{}

func (acr *BookSqlcRepository) Create(author model.BookInputModel) error {
	// TODO
	return nil
}

func (acr *BookSqlcRepository) Update(authorId uuid.UUID, author model.BookInputModel) error {
	// TODO
	return nil
}

func (acr *BookSqlcRepository) All() ([]model.Book, error) {
	// TODO
	return nil, nil
}

func (acr *BookSqlcRepository) Find(authorId uuid.UUID) (*model.Book, error) {
	// TODO
	return nil, nil
}

func (acr *BookSqlcRepository) Delete(authorId uuid.UUID) error {
	// TODO
	return nil
}

func (acr *BookSqlcRepository) Query(params map[string]any) ([]model.Book, error) {
	// TODO
	return nil, nil
}
