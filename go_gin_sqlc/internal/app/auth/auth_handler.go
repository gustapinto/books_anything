package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/domain/user"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/infrastructure/auth"
)

type AuthController struct {
	UserRepository user.UserRepository
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := ac.UserRepository.FindByUsernameAndPassword(credentials.Username, credentials.Password)
	if err != nil {
		if errors.Is(err, auth.ErrInvalidAuthentication) {
			c.Status(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	token, err := auth.GenerateToken(*user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
