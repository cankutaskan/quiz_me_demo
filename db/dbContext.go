package db

import (
	"math/rand"
	"sync"
	"time"

	"quiz_me/db/entities"
)

// DBContext represents the database context
type DBContext struct {
	questions map[int]entities.Question
	responses map[string]map[int]entities.Response
	stats     map[string]entities.Result
	mu        sync.RWMutex
	rng       *rand.Rand // Add a rand.Rand object for random number generation
}

// NewDBContext creates a new database context with a seeded random number generator
func NewDBContext() *DBContext {
	return &DBContext{
		questions: make(map[int]entities.Question),
		responses: make(map[string]map[int]entities.Response),
		stats:     make(map[string]entities.Result),
		rng:       rand.New(rand.NewSource(time.Now().UnixNano())), // Create a new rand.Rand object with a seed
	}
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

	if totalQuestionsToReturn <= 0 || len(db.questions) == 0 {
		return nil // Return nil if the requested number is non-positive or no questions are available
	}

	// Group questions by category
	categoryMap := make(map[string][]entities.Question)
	for _, question := range db.questions {
		categoryMap[question.Category] = append(categoryMap[question.Category], question)
	}

	// Adjust the number of questions to return if it's greater than the total available
	if totalQuestionsToReturn > len(db.questions) {
		totalQuestionsToReturn = len(db.questions)
	}

	var questionsToReturn []entities.Question
	selectedQuestions := make(map[int]struct{})

	// Calculate the number of questions to return from each category
	for _, questions := range categoryMap {
		categoryProportion := float64(len(questions)) / float64(len(db.questions))
		numQuestionsFromCategory := int(categoryProportion * float64(totalQuestionsToReturn))

		// Ensure at least one question is selected from each category if possible
		if numQuestionsFromCategory == 0 && len(questionsToReturn) < totalQuestionsToReturn {
			numQuestionsFromCategory = 1
		}

		// Select random questions from the category
		for i := 0; i < numQuestionsFromCategory && len(questionsToReturn) < totalQuestionsToReturn; i++ {
			for {
				index := db.rng.Intn(len(questions))
				question := questions[index]

				if _, exists := selectedQuestions[question.ID]; !exists {
					questionsToReturn = append(questionsToReturn, question)
					selectedQuestions[question.ID] = struct{}{}
					break
				}
			}
		}
	}

	return questionsToReturn
}

// AddResponse accepts a list of responses and resets previous answers for the given participant,
// recalculating their score based on the new responses.
func (db *DBContext) AddResponse(responses []entities.Response) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if len(responses) == 0 {
		return // No responses to process
	}

	participantID := responses[0].ParticipantID // Assuming all responses are from the same participant

	// Reset previous responses and stats for this participant
	db.responses[participantID] = make(map[int]entities.Response)
	db.stats[participantID] = entities.Result{TotalAnswers: len(responses)}

	// Process each response
	for _, response := range responses {
		db.responses[participantID][response.QuestionID] = response

		// Update the score if the answer is correct
		if db.isAnswerCorrect(response) {
			stats := db.stats[participantID]
			stats.CorrectAnswers++
			db.stats[participantID] = stats // Save the updated stats back to the map
		}
	}
}

// isAnswerCorrect checks if a given response is correct
func (db *DBContext) isAnswerCorrect(response entities.Response) bool {
	question, qExists := db.questions[response.QuestionID]
	if !qExists {
		return false
	}

	for _, answer := range question.Answers {
		if answer.ID == response.AnswerID {
			return answer.IsCorrect
		}
	}
	return false
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

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
