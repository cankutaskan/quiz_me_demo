package db

import (
	"sync"
)

// Answer represents an answer option for a question.
type Answer struct {
	Text      string
	IsCorrect bool
}

// Question represents a quiz question.
type Question struct {
	ID             int
	Text           string
	Answers        []Answer
	AlternativeIDs []int // IDs of alternative questions
}

// Response represents an answer submitted by a participant.
type Response struct {
	ParticipantID       string
	QuestionID          int
	SelectedAnswerIndex int // Index of the selected answer
}

// ParticipantStats holds statistics for a participant.
type ParticipantStats struct {
	ParticipantID  string
	CorrectAnswers int
	TotalAnswers   int
}

// InMemoryDB is a simple in-memory database for quiz questions and responses.
type InMemoryDB struct {
	questions map[int]Question
	responses []Response
	stats     map[string]ParticipantStats
	mu        sync.RWMutex
}

// NewInMemoryDB creates a new instance of InMemoryDB.
func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		questions: make(map[int]Question),
		responses: []Response{},
		stats:     make(map[string]ParticipantStats),
	}
}

// AddQuestion adds a new question to the database.
func (db *InMemoryDB) AddQuestion(q Question) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.questions[q.ID] = q
}

// GetAllQuestions retrieves all questions from the database.
func (db *InMemoryDB) GetAllQuestions() []Question {
	db.mu.RLock()
	defer db.mu.RUnlock()

	questions := make([]Question, 0, len(db.questions))
	for _, question := range db.questions {
		questions = append(questions, question)
	}
	return questions
}

// AddResponse adds a participant's response and updates their statistics.
func (db *InMemoryDB) AddResponse(r Response) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.responses = append(db.responses, r)
	question, exists := db.questions[r.QuestionID]
	if !exists {
		return
	}

	stats := db.stats[r.ParticipantID]
	stats.TotalAnswers++

	if question.Answers[r.SelectedAnswerIndex].IsCorrect {
		stats.CorrectAnswers++
	}

	db.stats[r.ParticipantID] = stats
}

// AddResponses adds multiple responses to the database.
func (db *InMemoryDB) AddResponses(responses []Response) {
	db.mu.Lock()
	defer db.mu.Unlock()

	for _, response := range responses {
		db.AddResponse(response)
	}
}

// GetParticipantStats retrieves statistics for a participant.
func (db *InMemoryDB) GetParticipantStats(participantID string) ParticipantStats {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.stats[participantID]
}

// CalculatePerformance calculates the performance percentage of a participant.
func (db *InMemoryDB) CalculatePerformance(participantID string) float64 {
	db.mu.RLock()
	defer db.mu.RUnlock()

	totalParticipants := len(db.stats)
	if totalParticipants == 0 {
		return 0
	}

	participantStats := db.stats[participantID]
	if participantStats.TotalAnswers == 0 {
		return 0
	}

	// Calculate how many participants performed better than the given participant
	betterThanCount := 0
	for _, stats := range db.stats {
		if stats.CorrectAnswers > participantStats.CorrectAnswers {
			betterThanCount++
		}
	}

	return float64(betterThanCount) / float64(totalParticipants) * 100
}
