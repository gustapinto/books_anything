package adapter

import (
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/model"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/sqlc/out"
)

func UserFromSqlcUser(sqlcUser *out.User, omitPassword bool) *model.User {
	password := sqlcUser.Password

	if omitPassword {
		password = ""
	}

	return &model.User{
		Id:        sqlcUser.ID,
		CreatedAt: sqlcUser.CreatedAt,
		UpdatedAt: sqlcUser.UpdatedAt,
		Name:      sqlcUser.Name,
		Username:  sqlcUser.Username,
		Password:  password,
	}
}

func UserFromSqlcCreateUserRow(sqlcUser *out.CreateUserRow) *model.User {
	return &model.User{
		Id:        sqlcUser.ID,
		CreatedAt: sqlcUser.CreatedAt,
		UpdatedAt: sqlcUser.UpdatedAt,
		Name:      sqlcUser.Name,
		Username:  sqlcUser.Username,
	}
}

func UserFromSqlcUpdateUserRow(sqlcUser *out.UpdateUserRow) *model.User {
	return &model.User{
		Id:        sqlcUser.ID,
		CreatedAt: sqlcUser.CreatedAt,
		UpdatedAt: sqlcUser.UpdatedAt,
		Name:      sqlcUser.Name,
		Username:  sqlcUser.Username,
	}
}

func AuthorFromSqlcAuthor(sqlcAuthor *out.Author) *model.Author {
	return &model.Author{
		Id:        sqlcAuthor.ID,
		CreatedAt: sqlcAuthor.CreatedAt,
		UpdatedAt: sqlcAuthor.UpdatedAt,
		Name:      sqlcAuthor.Name,
		UserId:    sqlcAuthor.UserID,
	}
}
