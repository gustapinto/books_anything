package controller

import (
	"net/http"
	"regexp"

	"github.com/gustapinto/books_rest/go_std/model"
	"github.com/gustapinto/books_rest/go_std/repository"
)

const (
	singleUserRoutePattern = `\/user\/([0-9]+)`
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
		if matched, _ := regexp.MatchString(singleUserRoutePattern, r.URL.Path); matched {
			usersController.Get(w, r)
		} else {
			usersController.GetAll(w, r)
		}
	case http.MethodPost:
		usersController.Create(w, r)
	case http.MethodPut:
		usersController.Update(w, r)
	case http.MethodDelete:
		usersController.Delete(w, r)
	default:
		MethodNotAllowed(w, r)
	}
}

func (usersController *UsersController) Get(w http.ResponseWriter, r *http.Request) {
	// TODO

	// userId, err := ExtractIdFromUrl(r.URL.Path)
	// if err != nil {
	// 	BadRequestResponse(w)
	// 	return
	// }
}

func (usersController *UsersController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := usersController.usersRepository.GetAll(true)
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

	if user.Id == 0 {
		userId, err := ExtractIdFromUrl(r.URL.Path)
		if err != nil {
			BadRequestResponse(w)
			return
		}

		user.Id = userId
	}

	newUser, err := usersController.usersRepository.Update(&user)
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
