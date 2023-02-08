package controller

import (
	"net/http"

	"github.com/gustapinto/books_rest/go_std/model"
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
	id, err := ExtractIdFromUrl(r.URL.Path)
	if err != nil {
		BadRequestResponse(w)
	}

	book, err := c.booksRepository.Find(id)
	if err != nil {
		ServerErrorResponse(w, err)
	}

	JsonResponse[model.Book](w, book, http.StatusOK)
}

func (c *BooksController) GetAll(w http.ResponseWriter, r *http.Request) {
	books, err := c.booksRepository.All()
	if err != nil {
		ServerErrorResponse(w, err)
		return
	}

	JsonResponse[[]model.Book](w, books, http.StatusOK)
}

func (c *BooksController) Create(w http.ResponseWriter, r *http.Request) {
	loggedUser := r.Context().Value("user").(model.User)
	book, err := UnmarshalJsonRequest[model.Book](w, r)
	if err != nil {
		BadRequestResponse(w)
		return
	}

	book.Creator.Id = loggedUser.Id

	newBook, err := c.booksRepository.Create(&book)
	if err != nil {
		ServerErrorResponse(w, err)
		return
	}

	JsonResponse[model.Book](w, newBook, http.StatusCreated)
}

func (c *BooksController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := ExtractIdFromUrl(r.URL.Path)
	if err != nil {
		BadRequestResponse(w)
		return
	}

	loggedUser := r.Context().Value("user").(model.User)
	book, err := UnmarshalJsonRequest[model.Book](w, r)
	if err != nil {
		BadRequestResponse(w)
		return
	}

	book.Creator.Id = loggedUser.Id

	newBook, err := c.booksRepository.Update(id, &book)
	if err != nil {
		ServerErrorResponse(w, err)
		return
	}

	JsonResponse[model.Book](w, newBook, http.StatusOK)
}

func (c *BooksController) Delete(w http.ResponseWriter, r *http.Request) {
	bookId, err := ExtractIdFromUrl(r.URL.Path)
	if err != nil {
		BadRequestResponse(w)
		return
	}

	if err := c.booksRepository.Delete(bookId); err != nil {
		ServerErrorResponse(w, err)
		return
	}
}
