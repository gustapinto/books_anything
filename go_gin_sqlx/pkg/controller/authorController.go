package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gustapinto/books_rest/go_gin_sqlx/pkg/model"
	"github.com/gustapinto/books_rest/go_gin_sqlx/pkg/repository"
)

type AuthorController struct {
	authorRepository repository.CrudRepository[model.Author, model.AuthorCreator]
}

func NewAuthorController(authorRepository repository.CrudRepository[model.Author, model.AuthorCreator]) *AuthorController {
	return &AuthorController{
		authorRepository: authorRepository,
	}
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
	authorId, err := strconv.Atoi(c.Param("authorId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	author, err := ac.authorRepository.Find(uint(authorId))
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
	authors, err := ac.authorRepository.All()
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
	var author model.Author
	if err := c.BindJSON(&author); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	loggedUser := c.Value("user").(model.User)
	author.CreatedBy = loggedUser.Id

	newAuthor, err := ac.authorRepository.Create(author)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, newAuthor)
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
	authorId, err := strconv.Atoi(c.Param("authorId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	var author model.Author
	if err := c.BindJSON(&author); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	newAuthor, err := ac.authorRepository.Update(uint(authorId), author)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, newAuthor)
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
	authorId, err := strconv.Atoi(c.Param("authorId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	if err := ac.authorRepository.Delete(uint(authorId)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.Status(http.StatusNoContent)
}
