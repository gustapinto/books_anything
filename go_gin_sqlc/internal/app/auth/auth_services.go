package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/domain/user"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/infrastructure/config"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrMissingAuthorizationHeader = errors.New(`missing Authorization header`)
	ErrMissingBearerKey           = errors.New(`missing Bearer prefix on token`)
	ErrInvalidToken               = errors.New("invalid token")
	ErrInvalidAuthentication      = errors.New("invalid authentication")
)

type UserClaims struct {
	User user.User
	jwt.RegisteredClaims
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePasswords(password, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}

	return true
}

func GenerateToken(user user.User) (string, error) {
	claims := &UserClaims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.APP_SECRET))
	if err != nil {
		return "", err
	}

	return "Bearer " + token, nil
}

func AuthenticateFromHeader(header http.Header) (*user.User, error) {
	authorization := header.Get("Authorization")
	if authorization == "" {
		return nil, ErrMissingAuthorizationHeader
	}

	if !strings.Contains(authorization, "Bearer") {
		return nil, ErrMissingBearerKey
	}

	authorization = strings.ReplaceAll(authorization, "Bearer ", "")

	token, err := jwt.ParseWithClaims(authorization, new(UserClaims), func(t *jwt.Token) (any, error) {
		return []byte(config.APP_SECRET), nil
	})
	if err != nil {
		return nil, ErrInvalidToken
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	user := token.Claims.(*UserClaims).User

	return &user, nil
}
