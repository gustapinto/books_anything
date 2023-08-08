package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/domain/user"
)

type UserHandler struct {
	UserRepository user.UserRepository
}

// @Summary Get user
// @Tags Users
// @Description Get a single user entry
// @Produce	json
// @Success	200	{object} User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Param id path uint true "The user id"
// @Router	/user/{id} [get]
func (uc *UserHandler) Find(c *gin.Context) {
	userId, err := uuid.FromBytes([]byte(c.Param("userId")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := uc.UserRepository.Find(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary	Get all users
// @Tags Users
// @Description	Get all users entries
// @Produce	json
// @Success	200	{array} User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Router	/user [get]
func (uc *UserHandler) All(c *gin.Context) {
	users, err := uc.UserRepository.All(1)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Summary	Create user
// @Tags Users
// @Description	Create a new user entry
// @Accept json
// @Produce	json
// @Success	201	{object} User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Param user body UserInputModel true "The user to be created"
// @Router	/user [post]
func (uc *UserHandler) Create(c *gin.Context) {
	var _user UserInputModel
	if err := c.BindJSON(&_user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if _, err := uc.UserRepository.Create(user.User{
		Name:     _user.Name,
		Username: _user.Username,
		Password: _user.Password,
	}); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary	Update user
// @Tags Users
// @Description	Update a user entry
// @Accept json
// @Produce	json
// @Success	200	{object} User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Param id path uint true "The user id"
// @Param user body UserInputModel true "The user to be updated"
// @Router	/user/{id} [put]
func (uc *UserHandler) Update(c *gin.Context) {
	userId, err := uuid.FromBytes([]byte(c.Param("userId")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var _user UserInputModel
	if err := c.BindJSON(&_user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if _, err := uc.UserRepository.Update(userId, user.User{
		Name:     _user.Name,
		Username: _user.Username,
		Password: _user.Password,
	}); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary	Delete user
// @Tags Users
// @Description	Delete a user entry
// @Produce	json
// @Success	204
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security Bearer
// @Param id path uint true "The user id"
// @Router	/user/{id} [delete]
func (uc *UserHandler) Delete(c *gin.Context) {
	userId, err := uuid.FromBytes([]byte(c.Param("userId")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := uc.UserRepository.Delete(userId); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.Status(http.StatusNoContent)
}
