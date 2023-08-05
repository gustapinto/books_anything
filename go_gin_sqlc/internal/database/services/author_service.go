package services

import (
	"github.com/google/uuid"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/database/repository"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/model"
)

type AuthorService struct {
	AuthorRepository repository.AuthorRepository
}

func (as *AuthorService) Create(author model.AuthorInputModel) (*model.Author, error) {
	return as.AuthorRepository.Create(author)
}

func (as *AuthorService) Update(authorId uuid.UUID, author model.AuthorInputModel) (*model.Author, error) {
	return as.AuthorRepository.Update(authorId, author)
}

func (as *AuthorService) Find(authorId uuid.UUID, userId uuid.UUID) (*model.Author, error) {
	return as.AuthorRepository.Find(authorId, userId)
}

func (as *AuthorService) All(userId uuid.UUID, page uint) ([]model.Author, error) {
	return as.AuthorRepository.All(userId, page)
}

func (as *AuthorService) Delete(authorId uuid.UUID, userId uuid.UUID) error {
	return as.AuthorRepository.Delete(authorId, userId)
}
