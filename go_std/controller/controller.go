package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type RestController interface {
	Get(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

func RestRouting(baseUrl string, c RestController, w http.ResponseWriter, r *http.Request) {
	hasId, _ := regexp.MatchString(`/`+baseUrl+`/([0-9]+)`, r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		if hasId {
			c.Get(w, r)
		} else {
			c.GetAll(w, r)
		}
	case http.MethodPost:
		c.Create(w, r)
	case http.MethodPut:
		c.Update(w, r)
	case http.MethodDelete:
		c.Delete(w, r)
	}
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, "Method "+r.Method+" not allowed")
}

func BadRequestResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, "Bad request")
}

func ServerErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, err.Error())
}

func NotFoundResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, err.Error())
}

func UnmarshalJsonRequest[T any | []any](w http.ResponseWriter, r *http.Request) (entity T, err error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return entity, err
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		return entity, err
	}

	return entity, nil
}

func JsonResponse[T any | []any](w http.ResponseWriter, entity any, statusCode int) {
	entityJson, err := json.Marshal(&entity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.WriteHeader(statusCode)
	w.Write(entityJson)
	return
}

func ExtractIdFromUrl(urlPath string) (uint, error) {
	splittedUrl := strings.Split(urlPath, "/")
	id, err := strconv.Atoi(splittedUrl[len(splittedUrl)-1])
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
