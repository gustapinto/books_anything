package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/model"
)

var ErrInvalidAuthentication = errors.New("invalid authentication")

type UserRepository interface {
	Create(model.UserInputModel) (*model.User, error)

	Update(uuid.UUID, model.UserInputModel) (*model.User, error)

	Find(uuid.UUID) (*model.User, error)

	All(uint) ([]model.User, error)

	Delete(uuid.UUID) error

	Login(string, string) (*model.User, error)
}

type AuthorRepository interface {
	Create(model.AuthorInputModel) (*model.Author, error)

	Update(uuid.UUID, model.AuthorInputModel) (*model.Author, error)

	Find(uuid.UUID, uuid.UUID) (*model.Author, error)

	All(uuid.UUID, uint) ([]model.Author, error)

	Delete(uuid.UUID, uuid.UUID) error
}

type BookRepository interface {
	Create(model.BookInputModel) (*model.Book, error)

	Update(uuid.UUID, model.BookInputModel) (*model.Book, error)

	Find(uuid.UUID, uuid.UUID) (*model.Book, error)

	All(uuid.UUID, uint) ([]model.Book, error)

	Delete(uuid.UUID, uuid.UUID) error
}
