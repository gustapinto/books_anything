package services

import (
	"github.com/google/uuid"
	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/database/repository"
	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/model"
)

type AuthorService struct {
	AuthorRepository repository.AuthorRepository
}

func (as *AuthorService) Create(author model.AuthorInputModel) error {
	return as.AuthorRepository.Create(author)
}

func (as *AuthorService) Update(authorId uuid.UUID, author model.AuthorInputModel) error {
	return as.AuthorRepository.Update(authorId, author)
}

func (as *AuthorService) Find(authorId uuid.UUID) (*model.AuthorViewModel, error) {
	return as.AuthorRepository.Find(authorId)
}

func (as *AuthorService) All() ([]model.AuthorViewModel, error) {
	return as.AuthorRepository.All()
}

func (as *AuthorService) Delete(authorId uuid.UUID) error {
	return as.AuthorRepository.Delete(authorId)
}

func (as *AuthorService) FindByName(authorName string) ([]model.AuthorViewModel, error) {
	return as.AuthorRepository.Query(map[string]any{
		"name": authorName,
	})
}
