package services

import (
	"github.com/google/uuid"
	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/database/repository"
	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/model"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (as *UserService) Create(author model.UserInputModel) error {
	return as.UserRepository.Create(author)
}

func (as *UserService) Update(authorId uuid.UUID, author model.UserInputModel) error {
	return as.UserRepository.Update(authorId, author)
}

func (as *UserService) Find(authorId uuid.UUID) (*model.UserViewModel, error) {
	return as.UserRepository.Find(authorId)
}

func (as *UserService) All() ([]model.UserViewModel, error) {
	return as.UserRepository.All()
}

func (as *UserService) Delete(authorId uuid.UUID) error {
	return as.UserRepository.Delete(authorId)
}

func (as *UserService) Login(username, password string) (*model.UserViewModel, error) {
	return as.UserRepository.Login(username, password)
}
