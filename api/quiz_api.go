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

type QuizAPI struct {
	db *db.DBContext
}

func NewQuizService(db *db.DBContext) *QuizAPI {
	return &QuizAPI{db: db}
}

func (s *QuizAPI) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/quiz", s.getQuiz).Methods("GET")
	r.HandleFunc("/quiz/{questionCount:[0-9]+}", s.getQuiz).Methods("GET")
	r.HandleFunc("/quiz/responses", s.submitAnswers).Methods("POST")
	r.HandleFunc("/quiz/performance/{participantID}", s.getPerformance).Methods("GET")
}

func (s *QuizAPI) getQuiz(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	questionCountStr := vars["questionCount"]

	questionCount := utils.ParseQueryParam(questionCountStr, 10) // Default to 10 questions if not provided

	questions := s.db.GetRandomQuestions(questionCount) // Delegate to the DB method
	quiz := get.Convert(questions)

	utils.EncodeJSONResponse(w, http.StatusOK, quiz)
}

func (s *QuizAPI) submitAnswers(w http.ResponseWriter, r *http.Request) {
	var responses post.Responses
	if err := json.NewDecoder(r.Body).Decode(&responses); err != nil {
		utils.EncodeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		return
	}

	if !s.validateQuestions(responses, w) {
		return
	}

	var entityResponses []entities.Response

	for _, response := range responses.Responses {
		entityResponses = append(entityResponses, entities.Response{
			ParticipantID: responses.UserID,
			QuestionID:    response.QuestionID,
			AnswerID:      response.AnswerID,
		})
	}

	s.db.AddResponse(entityResponses)

	resultsEntity := s.db.GetResult(responses.UserID)

	resultsAPI := s.calculateResult(resultsEntity)
	utils.EncodeJSONResponse(w, http.StatusOK, resultsAPI)
}

func (s *QuizAPI) getPerformance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	participantID := vars["participantID"]

	participantPerformance, comparisonPercentage := s.db.CalculatePerformance(participantID)

	performanceModel := get.Performance{
		Performance:          participantPerformance,
		ComparisonPercentage: comparisonPercentage,
	}

	utils.EncodeJSONResponse(w, http.StatusOK, performanceModel)
}

func (s *QuizAPI) validateQuestions(responses post.Responses, w http.ResponseWriter) bool {
	questionIDMap := make(map[int]bool)
	for _, response := range responses.Responses {
		if _, exists := questionIDMap[response.QuestionID]; exists {
			utils.EncodeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Duplicate answer submission detected for question"})
			return false
		}
		questionIDMap[response.QuestionID] = true
	}
	return true
}

func (s *QuizAPI) calculateResult(resultEntity entities.Result) get.Result {
	percentage := 0.0
	if resultEntity.TotalAnswers > 0 {
		percentage = (float64(resultEntity.CorrectAnswers) / float64(resultEntity.TotalAnswers)) * 100
	}

	return get.Result{
		CorrectAnswers: resultEntity.CorrectAnswers,
		TotalAnswers:   resultEntity.TotalAnswers,
		Percentage:     percentage,
	}
}
