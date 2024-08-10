package api

import (
	"encoding/json"
	"net/http"

	"quiz_me/api/models/get"
	"quiz_me/api/models/post"
	"quiz_me/db"
	"quiz_me/db/entities"

	"github.com/gorilla/mux"
)

// QuizService represents the service for handling quiz-related operations
type QuizService struct {
	db *db.DBContext
}

// NewQuizService creates a new QuizService
func NewQuizService(db *db.DBContext) *QuizService {
	return &QuizService{db: db}
}

// RegisterRoutes registers the routes for the quiz service
func (s *QuizService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/questions", s.getQuiz).Methods("GET")
	r.HandleFunc("/answers", s.handlePostAnswers).Methods("POST")
	r.HandleFunc("/performance/{participantID}", s.handleGetPerformance).Methods("GET")
}

// getQuiz retrieves all questions and returns them in the API model format
func (s *QuizService) getQuiz(w http.ResponseWriter, r *http.Request) {
	questions := s.db.GetAllQuestions() // Delegate to the DB method
	quiz := get.TransformQuestions(questions)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(quiz); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

// handlePostAnswers processes the answers submitted by a participant
func (s *QuizService) handlePostAnswers(w http.ResponseWriter, r *http.Request) {
	var responses post.Responses
	if err := json.NewDecoder(r.Body).Decode(&responses); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Convert post.AnswerResponse to entities.Response
	for _, response := range responses.Responses {
		entityResponse := entities.Response{
			ParticipantID: responses.UserID,
			QuestionID:    response.QuestionID,
			AnswerID:      response.AnswerID,
		}
		s.db.AddResponse(entityResponse)
	}

	stats := s.db.GetParticipantStats(responses.UserID)
	response := map[string]int{
		"correctAnswers": stats.CorrectAnswers,
		"totalAnswers":   stats.TotalAnswers,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

// handleGetPerformance calculates and returns the performance of a participant
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
