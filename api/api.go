package api

import (
	"log"
	"net/http"

	"quiz_me/db"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *db.InMemoryDB
}

func NewAPIServer(addr string, db *db.InMemoryDB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Initialize and register the QuizService with the subrouter
	quizService := NewQuizService(s.db)
	quizService.RegisterRoutes(subrouter)

	log.Println("Starting the API server at", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, router)) // Ensure to use the router with prefix applied
}
