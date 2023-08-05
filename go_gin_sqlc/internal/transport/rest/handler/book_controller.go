package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/database/services"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/model"
)

type BookController struct {
	BookService *services.BookService
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
	bookId, err := uuid.FromBytes([]byte(c.Param("bookId")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	book, err := bc.BookService.Find(bookId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, book)
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
	books, err := bc.BookService.All()
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
	var book model.BookInputModel
	if err := c.BindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	if err := bc.BookService.Create(book); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.Status(http.StatusCreated)
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
	bookId, err := uuid.FromBytes([]byte(c.Param("bookId")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	var book model.BookInputModel
	if err := c.BindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	if err := bc.BookService.Update(bookId, book); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.Status(http.StatusOK)
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
	bookId, err := uuid.FromBytes([]byte(c.Param("bookId")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	if err := bc.BookService.Delete(bookId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.Status(http.StatusNoContent)
}
