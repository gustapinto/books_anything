package adapter

import (
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/model"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/sqlc/out"
)

func AuthorFromSqlcAuthor(sqlcAuthor *out.Author) *model.Author {
	return &model.Author{
		Id:        sqlcAuthor.ID,
		CreatedAt: sqlcAuthor.CreatedAt,
		UpdatedAt: sqlcAuthor.UpdatedAt,
		Name:      sqlcAuthor.Name,
		UserId:    sqlcAuthor.UserID,
	}
}
