package controller

import (
	"io"
	"net/http"
)

type PingController struct{}

func (c *PingController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.Pong(w, r)
	default:
		MethodNotAllowed(w, r)
	}
}

func (c *PingController) Pong(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Pong")
}
