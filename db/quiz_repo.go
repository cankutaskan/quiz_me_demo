package db

import (
	"quiz_me/db/entities"
)

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
	db.results[participantID] = entities.Result{TotalAnswers: len(responses)}

	for _, response := range responses {
		db.responses[participantID][response.QuestionID] = response

		if db.isAnswerCorrect(response) {
			results := db.results[participantID]
			results.CorrectAnswers++
			db.results[participantID] = results
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

func (db *DBContext) GetResult(participantID string) entities.Result {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.results[participantID]
}

func (db *DBContext) CalculatePerformance(participantID string) (float64, float64) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	participantResults := db.results[participantID]

	participantPerformance := 0.0
	if participantResults.TotalAnswers > 0 {
		participantPerformance = float64(participantResults.CorrectAnswers) / float64(participantResults.TotalAnswers) * 100
	}

	totalParticipants := 0
	betterCount := 0

	for _, result := range db.results {
		if result.TotalAnswers > 0 {
			totalParticipants++
			overallPerformance := float64(result.CorrectAnswers) / float64(result.TotalAnswers) * 100
			if overallPerformance < participantPerformance {
				betterCount++
			}
		}
	}

	comparisonPercentage := 0.0
	if totalParticipants > 0 {
		comparisonPercentage = float64(betterCount) / float64(totalParticipants) * 100
	}

	return participantPerformance, comparisonPercentage
}
