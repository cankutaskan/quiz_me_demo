package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	projectService := NewProjectService()
	projectService.RegisterRoutes(subrouter)

	log.Println("Starting the API server at", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
