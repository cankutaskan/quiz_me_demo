package api

import (
	"encoding/json"
	"net/http"
	"quiz_me/api/models/get"
	"quiz_me/api/models/post"
	"quiz_me/db"
	"quiz_me/db/entities"
	"quiz_me/utils"

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
	r.HandleFunc("/quiz", s.getQuiz).Methods("GET") // Route without questionCount
	r.HandleFunc("/quiz/{questionCount:[0-9]+}", s.getQuiz).Methods("GET")
	r.HandleFunc("/responses", s.handlePostAnswers).Methods("POST")
	r.HandleFunc("/performance/{participantID}", s.handleGetPerformance).Methods("GET")
}

// getQuiz retrieves a specified number of random questions and returns them in the API model format
func (s *QuizService) getQuiz(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	questionCountStr := vars["questionCount"]

	// Use parseQueryParam to get the number of questions to return, defaulting to 10 if not specified
	questionCount := utils.ParseQueryParam(questionCountStr, 10) // Default to 10 questions if not provided

	questions := s.db.GetRandomQuestions(questionCount) // Delegate to the DB method
	quiz := get.TransformQuestions(questions)

	utils.EncodeJSONResponse(w, http.StatusOK, quiz)
}

// handlePostAnswers processes the answers submitted by a participant
func (s *QuizService) handlePostAnswers(w http.ResponseWriter, r *http.Request) {
	var responses post.Responses
	if err := json.NewDecoder(r.Body).Decode(&responses); err != nil {
		utils.EncodeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		return
	}

	for _, response := range responses.Responses {
		entityResponse := entities.Response{
			ParticipantID: responses.UserID,
			QuestionID:    response.QuestionID,
			AnswerID:      response.AnswerID,
		}
		s.db.AddResponse(entityResponse)
	}

	stats := s.db.GetParticipantStats(responses.UserID)
	result := map[string]int{
		"correctAnswers": stats.CorrectAnswers,
		"totalAnswers":   stats.TotalAnswers,
	}

	utils.EncodeJSONResponse(w, http.StatusOK, result)
}

// handleGetPerformance calculates and returns the performance of a participant
func (s *QuizService) handleGetPerformance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	participantID := vars["participantID"]

	performance := s.db.CalculatePerformance(participantID) // Delegate to the DB method

	result := map[string]float64{
		"performance": performance,
	}

	utils.EncodeJSONResponse(w, http.StatusOK, result)
}
