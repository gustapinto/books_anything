package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/auth"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/database/services"
)

type AuthController struct {
	UserService *services.UserService
}

// @Summary	Auth
// @Tags Auth
// @Description	Authenticate a user, returning the JWT token
// @Accept	json
// @Produce	json
// @Success	200	{object} TokenResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401
// @Failure 500 {object} ErrorResponse
// @Param credentials body LoginRequest true "The user credentials"
// @Router	/auth [post]
func (ac *AuthController) Login(c *gin.Context) {
	var credentials LoginRequest

	if err := c.BindJSON(&credentials); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	user, err := ac.UserService.Login(credentials.Username, credentials.Password)
	if err != nil {
		// if errors.Is(err, repository.ErrInvalidAuthentication) {
		// c.Status(http.StatusUnauthorized)
		// return
		// }

		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	token, err := auth.GenerateToken(*user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, NewTokenResponse(token))
}
