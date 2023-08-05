package repository

import "github.com/google/uuid"

type CrudRepository[T, K any] interface {
	Create(T) error

	Update(uuid.UUID, T) error

	Find(uuid.UUID) (*K, error)

	All() ([]K, error)

	Delete(uuid.UUID) error
}

type QueryRepository[T any] interface {
	Query(map[string]any) ([]T, error)
}
