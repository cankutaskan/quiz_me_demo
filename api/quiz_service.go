package api

import (
	"encoding/json"
	"net/http"

	"quiz_me/db" // Import the db package

	"github.com/gorilla/mux"
)

type QuizService struct {
	db *db.InMemoryDB // Use the correct type from the db package
}

func NewQuizService(db *db.InMemoryDB) *QuizService {
	return &QuizService{db: db}
}

func (s *QuizService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/questions", s.handleGetQuestions).Methods("GET")
	r.HandleFunc("/answers", s.handlePostAnswers).Methods("POST")
	r.HandleFunc("/performance/{participantID}", s.handleGetPerformance).Methods("GET")
}

func (s *QuizService) handleGetQuestions(w http.ResponseWriter, r *http.Request) {
	questions := s.db.GetAllQuestions() // Delegate to the DB method

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(questions); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (s *QuizService) handlePostAnswers(w http.ResponseWriter, r *http.Request) {
	var answers []db.Response
	if err := json.NewDecoder(r.Body).Decode(&answers); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Call the method in the DB without worrying about locking
	s.db.AddResponses(answers)

	participantID := answers[0].ParticipantID
	stats := s.db.GetParticipantStats(participantID)
	response := map[string]int{
		"correctAnswers": stats.CorrectAnswers,
		"totalAnswers":   stats.TotalAnswers,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (s *QuizService) handleGetPerformance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	participantID := vars["participantID"]

	performance := s.db.CalculatePerformance(participantID) // Delegate to the DB method

	response := map[string]float64{
		"performance": performance,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
