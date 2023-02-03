package controller

import (
	"errors"
	"io"
	"net/http"

	"github.com/gustapinto/books_rest/go_std/auth"
	"github.com/gustapinto/books_rest/go_std/repository"
)

type AuthController struct {
	usersRepository *repository.UsersRepository
}

func NewAuthController(usersRepository *repository.UsersRepository) *AuthController {
	return &AuthController{
		usersRepository: usersRepository,
	}
}

func (authController *AuthController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		authController.Login(w, r)
	default:
		MethodNotAllowed(w, r)
	}
}

func (authController *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	credentials, err := UnmarshalJsonRequest[struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}](w, r)
	if err != nil {
		BadRequestResponse(w)
		return
	}

	user, err := authController.usersRepository.ValidateUser(credentials.Username, credentials.Password)
	if err != nil {
		if errors.Is(err, repository.ErrInvalidAuthentication) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ServerErrorResponse(w, err)
		return
	}

	token, err := auth.GenerateToken(user)
	if err != nil {
		ServerErrorResponse(w, err)
		return
	}

	io.WriteString(w, `{"token": "`+token+`"}`)
}
