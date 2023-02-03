package controller

import (
	"net/http"
	"regexp"

	"github.com/gustapinto/books_rest/go_std/middleware"
	"github.com/gustapinto/books_rest/go_std/model"
	"github.com/gustapinto/books_rest/go_std/repository"
)

const (
	GetUserPattern = `\/user\/([0-9]+)`
)

type UsersController struct {
	usersRepository *repository.UsersRepository
}

func NewUsersController(usersRepository *repository.UsersRepository) *UsersController {
	return &UsersController{
		usersRepository: usersRepository,
	}
}

func (usersController *UsersController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if matched, _ := regexp.MatchString(GetUserPattern, r.URL.Path); matched {
			middleware.AuthFunc(w, r, usersController.Get)
		} else {
			middleware.AuthFunc(w, r, usersController.GetAll)
		}
	case http.MethodPost:
		usersController.Create(w, r)
	case http.MethodPut:
		middleware.AuthFunc(w, r, usersController.Update)
	case http.MethodDelete:
		middleware.AuthFunc(w, r, usersController.Delete)
	default:
		MethodNotAllowed(w, r)
	}
}

func (usersController *UsersController) Get(w http.ResponseWriter, r *http.Request) {
	userId, err := ExtractIdFromUrl(r.URL.Path)
	if err != nil {
		BadRequestResponse(w)
		return
	}

	user, err := usersController.usersRepository.FindWithoutPassword(userId)
	if err != nil {
		ServerErrorResponse(w, err)
		return
	}

	JsonResponse[model.User](w, user, http.StatusOK)
}

func (usersController *UsersController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := usersController.usersRepository.AllWithoutPassword()
	if err != nil {
		ServerErrorResponse(w, err)
		return
	}

	JsonResponse[[]model.User](w, users, http.StatusOK)
}

func (usersController *UsersController) Create(w http.ResponseWriter, r *http.Request) {
	user, err := UnmarshalJsonRequest[model.User](w, r)
	if err != nil {
		BadRequestResponse(w)
		return
	}

	newUser, err := usersController.usersRepository.Create(&user)
	if err != nil {
		ServerErrorResponse(w, err)
		return
	}

	JsonResponse[model.User](w, newUser, http.StatusCreated)
}

func (usersController *UsersController) Update(w http.ResponseWriter, r *http.Request) {
	user, err := UnmarshalJsonRequest[model.User](w, r)
	if err != nil {
		BadRequestResponse(w)
		return
	}

	userId, err := ExtractIdFromUrl(r.URL.Path)
	if err != nil {
		BadRequestResponse(w)
		return
	}

	newUser, err := usersController.usersRepository.Update(userId, &user)
	if err != nil {
		ServerErrorResponse(w, err)
		return
	}

	JsonResponse[model.User](w, newUser, http.StatusOK)
}

func (usersController *UsersController) Delete(w http.ResponseWriter, r *http.Request) {
	userId, err := ExtractIdFromUrl(r.URL.Path)
	if err != nil {
		BadRequestResponse(w)
	}

	if err := usersController.usersRepository.Delete(userId); err != nil {
		ServerErrorResponse(w, err)
	}
}
