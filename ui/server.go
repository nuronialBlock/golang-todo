package ui

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	get    string = "GET"
	post   string = "POST"
	delete string = "DELETE"
)

var router = mux.NewRouter()

type Server struct {
	router *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer() *Server {
	return &Server{
		router,
	}
}
