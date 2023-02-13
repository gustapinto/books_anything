package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gustapinto/books_rest/go_gin_sqlx/pkg/model"
	"github.com/gustapinto/books_rest/go_gin_sqlx/pkg/repository"
)

type UserController struct {
	userRepository *repository.UserRepository
}

func NewUserController(userRepository *repository.UserRepository) *UserController {
	return &UserController{
		userRepository: userRepository,
	}
}

// @Summary Get user
// @Tags Users
// @Description Get a single user entry
// @Produce	json
// @Success	200	{object} model.User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Param id path uint true "User ID"
// @Router	/user/{id} [get]
func (uc *UserController) Find(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	user, err := uc.userRepository.Find(uint(userId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary	Get all users
// @Tags Users
// @Description	Get all users entries
// @Produce	json
// @Success	200	{array} model.User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Router	/user [get]
func (uc *UserController) All(c *gin.Context) {
	users, err := uc.userRepository.All()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Summary	Create user
// @Tags Users
// @Description	Create a new user entry
// @Accept json
// @Produce	json
// @Success	201	{object} model.User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router	/user [post]
func (uc *UserController) Create(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	newUser, err := uc.userRepository.Create(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

// @Summary	Create user
// @Tags Users
// @Description	Update a user entry
// @Accept json
// @Produce	json
// @Success	200	{object} model.User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Param id path uint true "User ID"
// @Router	/user/{id} [put]
func (uc *UserController) Update(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	newUser, err := uc.userRepository.Update(uint(userId), user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, newUser)
}

// @Summary	Delete user
// @Tags Users
// @Description	Delete a user entry
// @Produce	json
// @Success	204
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Param id path uint true "User ID"
// @Router	/user/{id} [delete]
func (uc *UserController) Delete(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	if err := uc.userRepository.Delete(uint(userId)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	c.Status(http.StatusNoContent)
}
