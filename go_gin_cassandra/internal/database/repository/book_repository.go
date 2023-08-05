package repository

import (
	"github.com/google/uuid"
	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/model"
)

type BookRepository interface {
	CrudRepository[model.BookInputModel, model.BookViewModel]
	QueryRepository[model.BookViewModel]
}

type BookCassandraRepository struct{}

func (acr *BookCassandraRepository) Create(author model.BookInputModel) error {
	// TODO
	return nil
}

func (acr *BookCassandraRepository) Update(authorId uuid.UUID, author model.BookInputModel) error {
	// TODO
	return nil
}

func (acr *BookCassandraRepository) All() ([]model.BookViewModel, error) {
	// TODO
	return nil, nil
}

func (acr *BookCassandraRepository) Find(authorId uuid.UUID) (*model.BookViewModel, error) {
	// TODO
	return nil, nil
}

func (acr *BookCassandraRepository) Delete(authorId uuid.UUID) error {
	// TODO
	return nil
}

func (acr *BookCassandraRepository) Query(params map[string]any) ([]model.BookViewModel, error) {
	// TODO
	return nil, nil
}
