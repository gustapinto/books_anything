package controller

import (
	"io"
	"net/http"
)

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, "Method "+r.Method+" not allowed")
}
