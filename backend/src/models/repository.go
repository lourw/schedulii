package models

type Repository[T any] interface {
	Create(entity T) error
	Read(entity T) (T, error)
	Update(entity T) error
}
