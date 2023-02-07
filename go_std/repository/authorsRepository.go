package repository

import (
	"database/sql"
	"time"

	"github.com/gustapinto/books_rest/go_std/model"
)

type AuthorsRepository struct {
	db              *sql.DB
	usersRepository *UsersRepository
}

func NewAuthorsRepository(db *sql.DB, usersRepository *UsersRepository) *AuthorsRepository {
	return &AuthorsRepository{
		db:              db,
		usersRepository: usersRepository,
	}
}

func (_ *AuthorsRepository) Model() model.ModelInterface {
	return new(model.Author)
}

func (r *AuthorsRepository) Find(authorId uint) (author model.Author, err error) {
	query := `SELECT * FROM "` + r.Model().Table() + `" WHERE "id" = $1`
	dest := []any{
		&author.Id,
		&author.Name,
		&author.Creator.Id,
		&author.CreatedAt,
		&author.UpdatedAt,
	}

	if err := r.db.QueryRow(query, authorId).Scan(dest...); err != nil {
		return author, err
	}

	author.Creator, err = r.usersRepository.FindWithoutPassword(author.Creator.Id)
	if err != nil {
		return
	}

	return author, nil
}

func (r *AuthorsRepository) All() ([]model.Author, error) {
	query := `SELECT * FROM "` + r.Model().Table() + `"`
	authors := make([]model.Author, 0)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var author model.Author
		dest := []any{
			&author.Id,
			&author.Name,
			&author.Creator.Id,
			&author.CreatedAt,
			&author.UpdatedAt,
		}

		if err := rows.Scan(dest...); err != nil {
			return nil, err
		}

		author.Creator, err = r.usersRepository.FindWithoutPassword(author.Creator.Id)
		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (r *AuthorsRepository) Create(author *model.Author) (model.Author, error) {
	query := `INSERT INTO "` + r.Model().Table() + `" ("name", "creator_id", "created_at", "updated_at")
		VALUES ($1, $2, $3, $4)
		RETURNING id;`

	newAuthor := *author
	newAuthor.CreatedAt = time.Now()
	newAuthor.UpdatedAt = time.Now()

	if err := r.db.QueryRow(query, newAuthor.Name, newAuthor.Creator.Id, newAuthor.CreatedAt, newAuthor.UpdatedAt).Scan(&newAuthor.Id); err != nil {
		return model.Author{}, err
	}

	return newAuthor, nil
}

func (r *AuthorsRepository) Update(authorId uint, author *model.Author) (newAuthor model.Author, err error) {
	newAuthor = *author
	newAuthor.Id = authorId
	newAuthor.UpdatedAt = time.Now()

	query := `UPDATE "` + r.Model().Table() + `"
		SET "name" = $1, "updated_at" = $2
		WHERE "id" = $3`
	if _, err = r.db.Exec(query, newAuthor.Name, newAuthor.UpdatedAt, newAuthor.Id); err != nil {
		return
	}

	author.Creator, err = r.usersRepository.FindWithoutPassword(newAuthor.Creator.Id)
	if err != nil {
		return
	}

	return newAuthor, nil
}

func (r *AuthorsRepository) Delete(authorId uint) (err error) {
	query := `DELETE FROM "` + r.Model().Table() + `" WHERE "id" = $1`

	if _, err := r.db.Exec(query, authorId); err != nil {
		return err
	}

	return nil
}

func (r *AuthorsRepository) AllWithCreator(creatorId uint) ([]model.Author, error) {
	authors, err := r.All()
	if err != nil {
		return nil, err
	}

	authorsForCreator := make([]model.Author, 0)
	for _, author := range authors {
		if author.Creator.Id != creatorId {
			continue
		}

		authorsForCreator = append(authorsForCreator, author)
	}

	return authorsForCreator, nil
}
