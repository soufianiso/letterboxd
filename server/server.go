package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soufianiso/letterboxd/utils"
	"github.com/soufianiso/letterboxd/handlers"
)



type Server struct{
	port  string
}


func NewServer(port string) *Server{
	return &Server{port: port }
}


func (s *Server) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/", utils.MiddlewearApi(handlers.GetFilms))
	http.ListenAndServe(s.port, router)
}


