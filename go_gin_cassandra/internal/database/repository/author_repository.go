package repository

import (
	"github.com/google/uuid"
	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/model"
)

type AuthorRepository interface {
	CrudRepository[model.AuthorInputModel, model.AuthorViewModel]
	QueryRepository[model.AuthorViewModel]
}

type AuthorCassandraRepository struct{}

func (acr *AuthorCassandraRepository) Create(author model.AuthorInputModel) error {
	// TODO
	return nil
}

func (acr *AuthorCassandraRepository) Update(authorId uuid.UUID, author model.AuthorInputModel) error {
	// TODO
	return nil
}

func (acr *AuthorCassandraRepository) All() ([]model.AuthorViewModel, error) {
	// TODO
	return nil, nil
}

func (acr *AuthorCassandraRepository) Find(authorId uuid.UUID) (*model.AuthorViewModel, error) {
	// TODO
	return nil, nil
}

func (acr *AuthorCassandraRepository) Delete(authorId uuid.UUID) error {
	// TODO
	return nil
}

func (acr *AuthorCassandraRepository) Query(params map[string]any) ([]model.AuthorViewModel, error) {
	// TODO
	return nil, nil
}
