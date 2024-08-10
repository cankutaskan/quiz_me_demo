package db

import (
	"math/rand"
	"sync"
	"time"

	"quiz_me/db/entities"
)

type DBContext struct {
	questions map[int]entities.Question
	responses map[string]map[int]entities.Response
	stats     map[string]entities.Result
	mu        sync.RWMutex
	rng       *rand.Rand
}

func NewDBContext() *DBContext {
	return &DBContext{
		questions: make(map[int]entities.Question),
		responses: make(map[string]map[int]entities.Response),
		stats:     make(map[string]entities.Result),
		rng:       rand.New(rand.NewSource(time.Now().UnixNano())), // Create a new rand.Rand object with a seed
	}
}

func (db *DBContext) AddQuestion(question entities.Question) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.questions[question.ID] = question
}

func (db *DBContext) GetRandomQuestions(totalQuestionsToReturn int) []entities.Question {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if totalQuestionsToReturn <= 0 || len(db.questions) == 0 {
		return nil
	}

	if totalQuestionsToReturn > len(db.questions) {
		totalQuestionsToReturn = len(db.questions)
	}

	var questionsToReturn []entities.Question
	selectedQuestions := make(map[int]struct{})

	allQuestions := make([]entities.Question, 0, len(db.questions))
	for _, question := range db.questions {
		allQuestions = append(allQuestions, question)
	}

	for len(questionsToReturn) < totalQuestionsToReturn {
		index := db.rng.Intn(len(allQuestions))
		question := allQuestions[index]

		if _, exists := selectedQuestions[question.ID]; !exists {
			questionsToReturn = append(questionsToReturn, question)
			selectedQuestions[question.ID] = struct{}{}
		}
	}

	return questionsToReturn
}

func (db *DBContext) AddResponse(responses []entities.Response) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if len(responses) == 0 {
		return
	}

	participantID := responses[0].ParticipantID

	db.responses[participantID] = make(map[int]entities.Response)
	db.stats[participantID] = entities.Result{TotalAnswers: len(responses)}

	for _, response := range responses {
		db.responses[participantID][response.QuestionID] = response

		if db.isAnswerCorrect(response) {
			stats := db.stats[participantID]
			stats.CorrectAnswers++
			db.stats[participantID] = stats // Save the updated stats back to the map
		}
	}
}

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

func (db *DBContext) GetParticipantStats(participantID string) entities.Result {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.stats[participantID]
}

func (db *DBContext) CalculatePerformance(participantID string) (float64, float64) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	// Get the stats for the current participant
	participantStats := db.stats[participantID]

	// Calculate the performance percentage for the current participant
	participantPerformance := 0.0
	if participantStats.TotalAnswers > 0 {
		participantPerformance = float64(participantStats.CorrectAnswers) / float64(participantStats.TotalAnswers) * 100
	}

	// Calculate the performance for all participants
	totalParticipants := 0
	betterCount := 0

	for _, stat := range db.stats {
		if stat.TotalAnswers > 0 {
			totalParticipants++
			overallPerformance := float64(stat.CorrectAnswers) / float64(stat.TotalAnswers) * 100
			if overallPerformance < participantPerformance {
				betterCount++
			}
		}
	}

	// Calculate the percentage of participants that the current participant outperformed
	comparisonPercentage := 0.0
	if totalParticipants > 0 {
		comparisonPercentage = float64(betterCount) / float64(totalParticipants) * 100
	}

	return participantPerformance, comparisonPercentage
}
