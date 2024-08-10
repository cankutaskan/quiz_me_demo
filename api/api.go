package api

import (
	"log"
	"net/http"

	"quiz_me/db"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *db.DBContext
}

func NewAPIServer(addr string, db *db.DBContext) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/quiz-me").Subrouter()

	quizService := NewQuizService(s.db)
	quizService.RegisterRoutes(subrouter)

	log.Println("Starting the API server at", s.addr)
	http.ListenAndServe(s.addr, router)
}
