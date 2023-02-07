package controller

import (
	"net/http"

	"github.com/gustapinto/books_rest/go_std/repository"
)

type BooksController struct {
	booksRepository *repository.BooksRepository
}

func NewBooksController(booksRepository *repository.BooksRepository) *BooksController {
	return &BooksController{
		booksRepository: booksRepository,
	}
}

func (c *BooksController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	CrudRouting("book", c, w, r)
}

func (c *BooksController) Get(w http.ResponseWriter, r *http.Request) {
	// TODO
	return
}

func (c *BooksController) GetAll(w http.ResponseWriter, r *http.Request) {
	// TODO
	return
}

func (c *BooksController) Create(w http.ResponseWriter, r *http.Request) {
	// TODO
	return
}

func (c *BooksController) Update(w http.ResponseWriter, r *http.Request) {
	// TODO
	return
}

func (c *BooksController) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO
	return
}
