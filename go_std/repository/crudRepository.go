package repository

import (
	"github.com/gustapinto/books_rest/go_std/model"
)

type Repository interface {
	Model() model.ModelInterface
}

type CrudRepository[T model.Migrator] interface {
	Find(uint) (T, error)
	All() ([]T, error)
	Create(*T) (T, error)
	Update(uint, *T) (T, error)
	Delete(uint) error
}
