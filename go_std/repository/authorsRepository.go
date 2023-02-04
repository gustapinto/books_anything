package repository

import (
	"database/sql"
	"time"

	"github.com/gustapinto/books_rest/go_std/model"
)

type AuthorsRepository struct {
	db *sql.DB
}

func NewAuthorsRepository(db *sql.DB) *AuthorsRepository {
	return &AuthorsRepository{
		db: db,
	}
}

func (_ *AuthorsRepository) Model() model.ModelInterface {
	return new(model.Author)
}

func (r *AuthorsRepository) Find(authorId uint) (author model.Author, err error) {
	query := `SELECT * FROM "` + r.Model().Table() + `" WHERE "id" = $1`

	if err := r.db.QueryRow(query, authorId).Scan(&author.Id, &author.Name, &author.CreatedAt, &author.UpdatedAt); err != nil {
		return author, err
	}

	return author, nil
}

func (r *AuthorsRepository) All() (authors []model.Author, err error) {
	// TODO
	return
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
	// TODO
	return
}

func (r *AuthorsRepository) Delete(authorId uint) (err error) {
	// TODO
	return
}
