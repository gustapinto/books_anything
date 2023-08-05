package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/adapter"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/model"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/sqlc/out"
)

type AuthorSqlcRepository struct {
	Queries out.Querier
}

func (acr *AuthorSqlcRepository) Create(author model.AuthorInputModel) (*model.Author, error) {
	outAuthor, err := acr.Queries.CreateAuthor(context.Background(), out.CreateAuthorParams{
		Name: author.Name,
	})
	if err != nil {
		return nil, err
	}

	return adapter.AuthorFromSqlcAuthor(&outAuthor), nil
}

func (acr *AuthorSqlcRepository) Update(authorId uuid.UUID, author model.AuthorInputModel) (*model.Author, error) {
	outAuthor, err := acr.Queries.UpdateAuthor(context.Background(), out.UpdateAuthorParams{
		ID:   authorId,
		Name: author.Name,
	})
	if err != nil {
		return nil, err
	}

	return adapter.AuthorFromSqlcAuthor(&outAuthor), nil
}

func (acr *AuthorSqlcRepository) All(userId uuid.UUID, page uint) ([]model.Author, error) {
	outAuthors, err := acr.Queries.AllAuthorsFromUser(context.Background(), out.AllAuthorsFromUserParams{
		UserID: userId,
		Page:   int32(page),
	})
	if err != nil {
		return nil, err
	}

	if outAuthors == nil {
		return []model.Author{}, nil
	}

	authors := make([]model.Author, len(outAuthors))

	for _, outAuthor := range outAuthors {
		authors = append(authors, *adapter.AuthorFromSqlcAuthor(&outAuthor))
	}

	return authors, nil
}

func (acr *AuthorSqlcRepository) Find(authorId uuid.UUID, userId uuid.UUID) (*model.Author, error) {
	outAuthor, err := acr.Queries.FindAuthorByIdAndUser(context.Background(), out.FindAuthorByIdAndUserParams{
		ID:     authorId,
		UserID: userId,
	})
	if err != nil {
		return nil, err
	}

	return adapter.AuthorFromSqlcAuthor(&outAuthor), nil
}

func (acr *AuthorSqlcRepository) Delete(authorId uuid.UUID, userId uuid.UUID) error {
	return acr.Queries.DeleteAuthor(context.Background(), out.DeleteAuthorParams{
		ID:     authorId,
		UserID: userId,
	})
}
