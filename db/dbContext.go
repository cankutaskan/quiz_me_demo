package db

import (
	"sync"

	"quiz_me/db/entities"
)

// DBContext is a simple in-memory database for quiz questions and responses
type DBContext struct {
	questions map[int]entities.Question
	responses []entities.Response
	stats     map[string]entities.Result
	mu        sync.RWMutex
}

// NewDBContext creates a new in-memory database
func NewDBContext() *DBContext {
	return &DBContext{
		questions: make(map[int]entities.Question),
		responses: []entities.Response{},
		stats:     make(map[string]entities.Result),
	}
}

// AddQuestion adds a new question to the database
func (db *DBContext) AddQuestion(question entities.Question) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.questions[question.ID] = question
}

// GetAllQuestions retrieves all questions from the database
func (db *DBContext) GetAllQuestions() []entities.Question {
	db.mu.RLock()
	defer db.mu.RUnlock()
	var questions []entities.Question
	for _, question := range db.questions {
		questions = append(questions, question)
	}
	return questions
}

// AddResponse adds a participant's response to a question
func (db *DBContext) AddResponse(response entities.Response) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.responses = append(db.responses, response)

	// Update participant stats
	stats, exists := db.stats[response.ParticipantID]
	if !exists {
		stats = entities.Result{}
	}

	question, qExists := db.questions[response.QuestionID]
	if !qExists {
		return
	}

	var isCorrect bool
	for _, answer := range question.Answers {
		if answer.ID == response.AnswerID {
			isCorrect = answer.IsCorrect
			break
		}
	}

	stats.TotalAnswers++
	if isCorrect {
		stats.CorrectAnswers++
	}
	db.stats[response.ParticipantID] = stats
}

// GetParticipantStats retrieves statistics for a participant
func (db *DBContext) GetParticipantStats(participantID string) entities.Result {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.stats[participantID]
}

// CalculatePerformance calculates the performance of a participant as a percentage
func (db *DBContext) CalculatePerformance(participantID string) float64 {
	db.mu.RLock()
	defer db.mu.RUnlock()

	stats := db.stats[participantID]
	total := 0
	correct := 0

	for _, stat := range db.stats {
		total += stat.TotalAnswers
		correct += stat.CorrectAnswers
	}

	if total == 0 {
		return 0
	}

	return float64(stats.CorrectAnswers) / float64(stats.TotalAnswers) * 100
}
