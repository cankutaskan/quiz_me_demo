package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ProjectService struct {
}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (s *ProjectService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/projects/{id}", s.handleGetProject).Methods("GET")
}

func (s *ProjectService) handleGetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Print the project ID to the console
	fmt.Printf("Project ID: %s\n", id)

	// Respond with a simple message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Project ID printed to console"))
}
