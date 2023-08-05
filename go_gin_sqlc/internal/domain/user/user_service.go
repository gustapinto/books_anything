package user

import (
	"github.com/google/uuid"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/database/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (as *UserService) Create(author UserInputModel) (*User, error) {
	return as.UserRepository.Create(author)
}

func (as *UserService) Update(authorId uuid.UUID, author UserInputModel) (*User, error) {
	return as.UserRepository.Update(authorId, author)
}

func (as *UserService) Find(authorId uuid.UUID) (*User, error) {
	return as.UserRepository.Find(authorId)
}

func (as *UserService) All(page uint) ([]User, error) {
	return as.UserRepository.All(page)
}

func (as *UserService) Delete(authorId uuid.UUID) error {
	return as.UserRepository.Delete(authorId)
}

func (as *UserService) Login(username, password string) (*User, error) {
	return as.UserRepository.Login(username, password)
}
