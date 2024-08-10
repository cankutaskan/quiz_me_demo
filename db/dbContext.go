package db

import (
	"math/rand"
	"sync"
	"time"

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

func init() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
}

// AddQuestion adds a new question to the database
func (db *DBContext) AddQuestion(question entities.Question) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.questions[question.ID] = question
}

// GetRandomQuestions retrieves a specified number of random questions proportionally from each category
func (db *DBContext) GetRandomQuestions(totalQuestionsToReturn int) []entities.Question {
	db.mu.RLock()
	defer db.mu.RUnlock()

	// Validate the input
	if totalQuestionsToReturn <= 0 {
		return nil // Return nil if the requested number is non-positive
	}

	// Group questions by category
	categoryMap := make(map[string][]entities.Question)
	for _, question := range db.questions {
		categoryMap[question.Category] = append(categoryMap[question.Category], question)
	}

	// Calculate total number of questions
	totalQuestions := len(db.questions)

	// Adjust the number of questions to return if it's greater than the total available
	if totalQuestionsToReturn > totalQuestions {
		totalQuestionsToReturn = totalQuestions
	}

	// Use a slice to keep track of unique questions
	questionsToReturn := make([]entities.Question, 0, totalQuestionsToReturn)
	selectedQuestions := make(map[int]struct{})

	// Calculate the number of questions to return from each category
	for _, questions := range categoryMap {
		numQuestions := int(float64(len(questions)) / float64(totalQuestions) * float64(totalQuestionsToReturn))
		if numQuestions > len(questions) {
			numQuestions = len(questions)
		}

		// Select random questions without shuffling
		for len(questionsToReturn) < totalQuestionsToReturn && numQuestions > 0 {
			index := rand.Intn(len(questions))
			question := questions[index]

			if _, exists := selectedQuestions[question.ID]; !exists {
				questionsToReturn = append(questionsToReturn, question)
				selectedQuestions[question.ID] = struct{}{}
				numQuestions--
			}
		}
	}

	return questionsToReturn
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
