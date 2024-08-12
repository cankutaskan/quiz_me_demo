package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"quiz_me/api/models/get"
	"quiz_me/api/models/post"
	"quiz_me/db"
	"quiz_me/db/entities"
	"quiz_me/utils"

	"github.com/gorilla/mux"
)

type QuizAPI struct {
	db   *db.DBContext
	quiz *get.Quiz
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

	questionCount := utils.ParseQueryParam(questionCountStr, 10)

	questions := s.db.GetRandomQuestions(questionCount)
	quiz := get.Convert(questions)

	s.quiz = &quiz

	utils.EncodeJSONResponse(w, http.StatusOK, quiz)
}

func (s *QuizAPI) submitAnswers(w http.ResponseWriter, r *http.Request) {
	var responses post.Responses
	if err := json.NewDecoder(r.Body).Decode(&responses); err != nil {
		utils.EncodeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		return
	}

	if s.quiz == nil {
		utils.EncodeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "No quiz available. Please fetch a quiz first."})
		return
	}

	if !s.validateQuestions(responses, w) {
		return
	}

	responseMap := createResponseModel(responses)

	entityResponses := prepareEntityResponses(responses.UserID, responseMap, s.quiz)

	s.db.AddResponse(entityResponses)

	resultsEntity := s.db.GetResult(responses.UserID)
	resultsAPI := s.calculateResult(resultsEntity)

	utils.EncodeJSONResponse(w, http.StatusOK, resultsAPI)
}

func (s *QuizAPI) getPerformance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	participantID := vars["participantID"]

	participantPerformance, comparisonPercentage := s.db.CalculatePerformance(participantID)

	message := fmt.Sprintf("You were better than %.2f%% of all quizzers", comparisonPercentage)

	performanceResponse := get.PerformanceResponse{
		Performance:          participantPerformance,
		ComparisonMessage:    message,
		ComparisonPercentage: comparisonPercentage,
	}

	utils.EncodeJSONResponse(w, http.StatusOK, performanceResponse)
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

func createResponseModel(responses post.Responses) map[int]int {
	responseMap := make(map[int]int)
	for _, response := range responses.Responses {
		responseMap[response.QuestionID] = response.AnswerID
	}
	return responseMap
}

func prepareEntityResponses(userID string, responseMap map[int]int, quiz *get.Quiz) []entities.Response {
	var entityResponses []entities.Response
	for _, question := range quiz.Questions {
		answerID, answered := responseMap[question.ID]
		if !answered {
			answerID = -1
		}
		entityResponses = append(entityResponses, entities.Response{
			ParticipantID: userID,
			QuestionID:    question.ID,
			AnswerID:      answerID,
		})
	}
	return entityResponses
}
