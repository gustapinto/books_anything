package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustapinto/books_rest/go_gin_sqlx/pkg/auth"
)

func Auth(c *gin.Context) {
	loggedUser, err := auth.AuthenticateFromHeader(c.Request.Header)
	if err != nil {
		if errors.Is(err, auth.ErrMissingAuthorizationHeader) || errors.Is(err, auth.ErrMissingBearerKey) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if errors.Is(err, auth.ErrInvalidToken) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Set("user", loggedUser)
	c.Next()
}
