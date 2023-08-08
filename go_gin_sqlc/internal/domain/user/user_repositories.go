package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/auth"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/sqlc/out"
)

type UserRepository interface {
	Create(user UserInputModel) (*User, error)

	Update(uuid.UUID, UserInputModel) (*User, error)

	Find(uuid.UUID) (*User, error)

	All(uint) (*UserPagination, error)

	Delete(uuid.UUID) error

	Login(string, string) (*User, error)
}

type UserSqlcRepository struct {
	Queries out.Querier
}

func (acr *UserSqlcRepository) Create(inUser UserInputModel) (*User, error) {
	hashedPassword, err := auth.HashPassword(inUser.Password)
	if err != nil {
		return nil, err
	}

	outUser, err := acr.Queries.CreateUser(context.Background(), out.CreateUserParams{
		Name:     inUser.Name,
		Username: inUser.Username,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}

	return UserFromSqlcCreateUserRow(&outUser), nil
}

func (acr *UserSqlcRepository) Update(userId uuid.UUID, inUser UserInputModel) (*User, error) {
	hashedPassword, err := auth.HashPassword(inUser.Password)
	if err != nil {
		return nil, err
	}

	outUser, err := acr.Queries.UpdateUser(context.Background(), out.UpdateUserParams{
		Name:     inUser.Name,
		Username: inUser.Username,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}

	return UserFromSqlcUpdateUserRow(&outUser), nil
}

func (acr *UserSqlcRepository) All(page uint) (*UserPagination, error) {
	outUsers, err := acr.Queries.AllUsers(context.Background(), int32(page))
	if err != nil {
		return nil, err
	}

	outUsersCount, err := acr.Queries.UsersCount(context.Background())
	if err != nil {
		return nil, err
	}

	if outUsers == nil {
		return &UserPagination{}, nil
	}

	users := make([]User, len(outUsers))

	for _, outUser := range outUsers {
		users = append(users, *UserFromSqlcUser(&outUser, true))
	}

	nextPage := page + 1

	if nextPage > uint(outUsersCount.TotalPages) {
		nextPage = 0
	}

	pagination := &UserPagination{
		Data:        users,
		TotalCount:  uint(outUsersCount.TotalCount),
		TotalPages:  uint(outUsersCount.TotalPages),
		CurrentPage: page,
		NextPage:    nextPage,
	}

	return pagination, nil
}

func (acr *UserSqlcRepository) Find(userId uuid.UUID) (*User, error) {
	outUser, err := acr.Queries.FindUserById(context.Background(), userId)
	if err != nil {
		return nil, err
	}

	return UserFromSqlcUser(&outUser, true), nil
}

func (acr *UserSqlcRepository) Delete(userId uuid.UUID) error {
	return acr.Queries.DeleteUser(context.Background(), userId)
}

func (acr *UserSqlcRepository) Login(username, password string) (*User, error) {
	outUser, err := acr.Queries.FindUserByUsername(context.Background(), username)
	if err != nil {
		return nil, err
	}

	if matched := auth.ComparePasswords(password, outUser.Password); !matched {
		return nil, auth.ErrInvalidAuthentication
	}

	return UserFromSqlcUser(&outUser, true), nil
}
