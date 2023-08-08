package user

import "github.com/gustapinto/books_rest/go_gin_sqlc_ddd/sqlc/out"

func UserFromSqlcUser(sqlcUser *out.User, omitPassword bool) *User {
	password := sqlcUser.Password

	if omitPassword {
		password = ""
	}

	return &User{
		Id:        sqlcUser.ID,
		CreatedAt: sqlcUser.CreatedAt,
		UpdatedAt: sqlcUser.UpdatedAt,
		Name:      sqlcUser.Name,
		Username:  sqlcUser.Username,
		Password:  password,
	}
}

func UserFromSqlcCreateUserRow(sqlcUser *out.CreateUserRow) *User {
	return &User{
		Id:        sqlcUser.ID,
		CreatedAt: sqlcUser.CreatedAt,
		UpdatedAt: sqlcUser.UpdatedAt,
		Name:      sqlcUser.Name,
		Username:  sqlcUser.Username,
	}
}

func UserFromSqlcUpdateUserRow(sqlcUser *out.UpdateUserRow) *User {
	return &User{
		Id:        sqlcUser.ID,
		CreatedAt: sqlcUser.CreatedAt,
		UpdatedAt: sqlcUser.UpdatedAt,
		Name:      sqlcUser.Name,
		Username:  sqlcUser.Username,
	}
}
