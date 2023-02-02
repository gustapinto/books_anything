package controller

import (
	"io"
	"net/http"
)

type PingController struct{}

func NewPingController() *PingController {
	return &PingController{}
}

func (pingController *PingController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		pingController.Pong(w, r)
	default:
		MethodNotAllowed(w, r)
	}
}

func (_ *PingController) Pong(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Pong")
}
