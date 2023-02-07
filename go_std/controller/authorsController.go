package controller

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gustapinto/books_rest/go_std/model"
	"github.com/gustapinto/books_rest/go_std/repository"
)

type AuthorsController struct {
	authorsRepository *repository.AuthorsRepository
}

func NewAuthorsController(authorsRepository *repository.AuthorsRepository) *AuthorsController {
	return &AuthorsController{
		authorsRepository: authorsRepository,
	}
}

func (c *AuthorsController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	CrudRouting("author", c, w, r)
}

func (c *AuthorsController) Get(w http.ResponseWriter, r *http.Request) {
	authorId, err := ExtractIdFromUrl(r.URL.Path)
	if err != nil {
		BadRequestResponse(w)
		return
	}

	author, err := c.authorsRepository.Find(authorId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			NotFoundResponse(w, err)
			return
		}

		ServerErrorResponse(w, err)
		return
	}

	JsonResponse[model.Author](w, author, http.StatusOK)
}

func (c *AuthorsController) GetAll(w http.ResponseWriter, r *http.Request) {
	loggedUser := r.Context().Value("user").(model.User)
	authors, err := c.authorsRepository.AllWithCreator(loggedUser.Id)
	if err != nil {
		ServerErrorResponse(w, err)
		return
	}

	JsonResponse[[]model.Author](w, authors, http.StatusOK)
}

func (c *AuthorsController) Create(w http.ResponseWriter, r *http.Request) {
	author, err := UnmarshalJsonRequest[model.Author](w, r)
	if err != nil {
		BadRequestResponse(w)
		return
	}

	loggedUser := r.Context().Value("user").(model.User)
	author.Creator = loggedUser

	newAuthor, err := c.authorsRepository.Create(&author)
	if err != nil {
		ServerErrorResponse(w, err)
	}

	JsonResponse[model.Author](w, newAuthor, http.StatusCreated)
}

func (c *AuthorsController) Update(w http.ResponseWriter, r *http.Request) {
	authorId, err := ExtractIdFromUrl(r.URL.Path)
	if err != nil {
		BadRequestResponse(w)
		return
	}

	author, err := UnmarshalJsonRequest[model.Author](w, r)
	if err != nil {
		BadRequestResponse(w)
		return
	}

	newAuthor, err := c.authorsRepository.Update(authorId, &author)
	if err != nil {
		ServerErrorResponse(w, err)
		return
	}

	JsonResponse[model.Author](w, newAuthor, http.StatusOK)
}

func (c *AuthorsController) Delete(w http.ResponseWriter, r *http.Request) {
	authorId, err := ExtractIdFromUrl(r.URL.Path)
	if err != nil {
		BadRequestResponse(w)
		return
	}

	if err := c.authorsRepository.Delete(authorId); err != nil {
		ServerErrorResponse(w, err)
		return
	}
}
