package repository

type CrudRepository[T, K any] interface {
	Find(uint) (*K, error)
	All() ([]K, error)
	Create(T) (*K, error)
	Update(uint, T) (*K, error)
	Delete(uint) error
}
