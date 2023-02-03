package repository

import (
	"database/sql"
	"time"

	"github.com/gustapinto/books_rest/go_std/model"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (usersRepository *UsersRepository) Model() model.ModelInterface {
	return new(model.User)
}

func (usersRepository *UsersRepository) Find(userId uint) (user model.User, err error) {
	query := `SELECT * FROM "` + usersRepository.Model().Table() + `" WHERE "id" = $1`

	if err := usersRepository.db.QueryRow(query, userId).Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return user, err
	}

	return user, nil
}

func (usersRepository *UsersRepository) All() ([]model.User, error) {
	users := make([]model.User, 0)

	query := `SELECT * FROM "` + usersRepository.Model().Table() + `"`

	rows, err := usersRepository.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.User

		if err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (usersRepository *UsersRepository) Create(user *model.User) (model.User, error) {
	newUser, err := usersRepository.PrepareUser(user)
	if err != nil {
		return model.User{}, err
	}

	query := `
		INSERT INTO "` + usersRepository.Model().Table() + `" ("name", "username", "password", "created_at", "updated_at")
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	if err := usersRepository.db.QueryRow(query, newUser.Name, newUser.Username, newUser.Password, newUser.CreatedAt, newUser.UpdatedAt).Scan(&newUser.Id); err != nil {
		return model.User{}, err
	}

	return newUser, nil
}

func (usersRepository *UsersRepository) Update(userId uint, user *model.User) (newUser model.User, err error) {
	newUser, err = usersRepository.PrepareUser(user)
	if err != nil {
		return
	}

	query := `
		UPDATE "` + usersRepository.Model().Table() + `"
		SET "name" = $1, "username" = $2, "password" = $3, "updated_at" = $4
		WHERE "id" = $5;
	`

	if _, err = usersRepository.db.Exec(query, newUser.Name, newUser.Username, newUser.Password, newUser.UpdatedAt, userId); err != nil {
		return
	}

	return newUser, nil
}

func (usersRepository *UsersRepository) Delete(userId uint) error {
	query := `
		DELETE FROM "` + usersRepository.Model().Table() + `"
		WHERE "id" = $1;
	`

	if _, err := usersRepository.db.Exec(query, userId); err != nil {
		return err
	}

	return nil
}

func (usersRepository *UsersRepository) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (usersRepository *UsersRepository) FindWithoutPassword(userId uint) (model.User, error) {
	user, err := usersRepository.Find(userId)
	if err != nil {
		return user, err
	}

	user.Password = ""

	return user, nil
}

func (usersRepository *UsersRepository) AllWithoutPassword() ([]model.User, error) {
	users, err := usersRepository.All()
	if err != nil {
		return nil, err
	}

	for i := range users {
		users[i].Password = ""
	}

	return users, nil
}

func (usersRepository *UsersRepository) PrepareUser(user *model.User) (model.User, error) {
	newUser := *user
	newUser.UpdatedAt = time.Now()

	if newUser.CreatedAt.IsZero() {
		newUser.CreatedAt = time.Now()
	}

	if newUser.Password != "" {
		hashedPassword, err := usersRepository.HashPassword(newUser.Password)
		if err != nil {
			return model.User{}, err
		}
		newUser.Password = hashedPassword
	}

	return newUser, nil
}
