package middleware

import (
	"errors"
	"io"
	"net/http"

	"github.com/gustapinto/books_rest/go_std/auth"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AuthFunc(w, r, next.ServeHTTP)
	})
}

func AuthFunc(w http.ResponseWriter, r *http.Request, next func(http.ResponseWriter, *http.Request)) {
	if err := auth.AuthenticateFromHeader(r.Header); err != nil {
		if errors.Is(err, auth.ErrMissingAuthorizationHeader) || errors.Is(err, auth.ErrMissingBearerKey) {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, err.Error())
			return
		}

		if errors.Is(err, auth.ErrInvalidToken) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	next(w, r)
}
