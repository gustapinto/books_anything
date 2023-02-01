package controller

import (
	"net/http"

	"github.com/gustapinto/books_rest/go_std/repository"
)

type UsersController struct {
	usersRepository *repository.UsersRepository
}

func NewUsersController(usersRepository *repository.UsersRepository) *UsersController {
	return &UsersController{
		usersRepository: usersRepository,
	}
}

func (c *UsersController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.Create(w, r)
	default:
		MethodNotAllowed(w, r)
	}
}

func (c *UsersController) Create(w http.ResponseWriter, r *http.Request) {
	/**
	TODO

	Add the user create method
	*/
}
