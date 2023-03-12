package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/database/repository"
	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/model"
)

type BookCrudRepository interface {
	repository.CrudRepository[model.Book, model.BookAuthorCreator]
}

type BookController struct {
	bookRepository BookCrudRepository
}

func NewBookController(bookRepository BookCrudRepository) *BookController {
	return &BookController{
		bookRepository: bookRepository,
	}
}

// @Summary Get book
// @Tags Books
// @Description Get a single book entry with relationships
// @Produce	json
// @Success	200	{object} model.BookAuthorCreator
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Param id path uint true "The book id"
// @Router	/book/{id} [get]
func (bc *BookController) Find(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Param("bookId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	author, err := bc.bookRepository.Find(uint(bookId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, author)
}

// @Summary Get all books
// @Tags Books
// @Description Get all books entries with relationships
// @Produce	json
// @Success	200	{array} model.BookAuthorCreator
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Router	/book [get]
func (bc *BookController) All(c *gin.Context) {
	books, err := bc.bookRepository.All()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, books)
}

// @Summary	Create book
// @Tags Books
// @Description	Create a new book entry, returning the created book with relationships
// @Accept json
// @Produce	json
// @Success	201	{object} model.BookAuthorCreator
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Param book body schema.NewBookPayload true "The book to be created"
// @Router	/book [post]
func (bc *BookController) Create(c *gin.Context) {
	var book model.Book
	if err := c.BindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	loggedUser := c.Value("user").(model.User)
	book.CreatedBy = loggedUser.Id

	newBook, err := bc.bookRepository.Create(book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, newBook)
}

// @Summary	Update book
// @Tags Books
// @Description	Update a book entry, returning the updated book with relationships
// @Accept json
// @Produce	json
// @Success	200	{object} model.BookAuthorCreator
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Param book body schema.NewBookPayload true "The book to be updated"
// @Param id path uint true "The book id"
// @Router	/book/{id} [put]
func (bc *BookController) Update(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Param("bookId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	var book model.Book
	if err := c.BindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	newBook, err := bc.bookRepository.Update(uint(bookId), book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, newBook)
}

// @Summary	Delete book
// @Tags Books
// @Description	Delete a book entry
// @Produce	json
// @Success	204
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Param id path uint true "The book id"
// @Router	/book/{id} [delete]
func (bc *BookController) Delete(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Param("bookId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	if err := bc.bookRepository.Delete(uint(bookId)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.Status(http.StatusNoContent)
}
