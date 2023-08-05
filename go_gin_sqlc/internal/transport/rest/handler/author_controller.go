package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/database/services"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/model"
)

type AuthorController struct {
	AuthorService *services.AuthorService
}

// @Summary Get author
// @Tags Authors
// @Description Get a single author entry with relationships
// @Produce	json
// @Success	200	{object} model.AuthorCreator
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Param id path uint true "The author id"
// @Router	/author/{id} [get]
func (ac *AuthorController) Find(c *gin.Context) {
	authorId, err := uuid.FromBytes([]byte(c.Param("authorId")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	author, err := ac.AuthorService.Find(authorId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, author)
}

// @Summary	Get all authors
// @Tags Authors
// @Description	Get all authors entries with relationships
// @Produce	json
// @Success	200	{array} model.Author
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Router	/author [get]
func (ac *AuthorController) All(c *gin.Context) {
	authors, err := ac.AuthorService.All()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, authors)
}

// @Summary	Create author
// @Tags Authors
// @Description	Create a new author entry
// @Accept json
// @Produce	json
// @Success	201	{object} model.AuthorCreator
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Param author body schema.NewAuthorPayload true "The author to be created"
// @Router	/author [post]
func (ac *AuthorController) Create(c *gin.Context) {
	var author model.AuthorInputModel
	if err := c.BindJSON(&author); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	if err := ac.AuthorService.Create(author); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary	Update author
// @Tags Authors
// @Description	Update a author entry
// @Accept json
// @Produce	json
// @Success	200	{object} model.AuthorCreator
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Param id path uint true "The author id"
// @Param author body schema.NewAuthorPayload true "The author to be updated"
// @Router	/author/{id} [put]
func (ac *AuthorController) Update(c *gin.Context) {
	authorId, err := uuid.FromBytes([]byte(c.Param("authorId")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	var author model.AuthorInputModel
	if err := c.BindJSON(&author); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	if err := ac.AuthorService.Update(authorId, author); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.Status(http.StatusOK)
}

// @Summary	Delete author
// @Tags Authors
// @Description	Delete a author entry
// @Produce	json
// @Success	204
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Param id path uint true "The author id"
// @Router	/author/{id} [delete]
func (ac *AuthorController) Delete(c *gin.Context) {
	authorId, err := uuid.FromBytes([]byte(c.Param("authorId")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	if err := ac.AuthorService.Delete(authorId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.Status(http.StatusNoContent)
}
