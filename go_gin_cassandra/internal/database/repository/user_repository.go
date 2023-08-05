package repository

import (
	"github.com/google/uuid"
	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/model"
)

type UserRepository interface {
	CrudRepository[model.UserInputModel, model.UserViewModel]
	QueryRepository[model.UserViewModel]

	Login(string, string) (*model.UserViewModel, error)
}

type UserCassandraRepository struct{}

func (acr *UserCassandraRepository) Create(author model.UserInputModel) error {
	// TODO
	return nil
}

func (acr *UserCassandraRepository) Update(authorId uuid.UUID, author model.UserInputModel) error {
	// TODO
	return nil
}

func (acr *UserCassandraRepository) All() ([]model.UserViewModel, error) {
	// TODO
	return nil, nil
}

func (acr *UserCassandraRepository) Find(authorId uuid.UUID) (*model.UserViewModel, error) {
	// TODO
	return nil, nil
}

func (acr *UserCassandraRepository) Delete(authorId uuid.UUID) error {
	// TODO
	return nil
}

func (acr *UserCassandraRepository) Query(params map[string]any) ([]model.UserViewModel, error) {
	// TODO
	return nil, nil
}

func (acr *UserCassandraRepository) Login(username, password string) (*model.UserViewModel, error) {
	// TODO
	return nil, nil
}
