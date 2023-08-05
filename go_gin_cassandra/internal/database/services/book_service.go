package services

import (
	"github.com/google/uuid"
	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/database/repository"
	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/model"
)

type BookService struct {
	BookRepository repository.BookRepository
}

func (as *BookService) Create(author model.BookInputModel) error {
	return as.BookRepository.Create(author)
}

func (as *BookService) Update(authorId uuid.UUID, author model.BookInputModel) error {
	return as.BookRepository.Update(authorId, author)
}

func (as *BookService) Find(authorId uuid.UUID) (*model.BookViewModel, error) {
	return as.BookRepository.Find(authorId)
}

func (as *BookService) All() ([]model.BookViewModel, error) {
	return as.BookRepository.All()
}

func (as *BookService) Delete(authorId uuid.UUID) error {
	return as.BookRepository.Delete(authorId)
}
