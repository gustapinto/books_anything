package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

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
