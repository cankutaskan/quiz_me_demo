package db

import (
	"quiz_me/db/entities"
	"sort"
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

		if isAnswerCorrect(db.questions, response) {
			results := db.results[participantID]
			results.CorrectAnswers++
			db.results[participantID] = results
		}
	}
}

func (db *DBContext) GetResult(participantID string) entities.Result {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.results[participantID]
}

func (db *DBContext) CalculatePerformance(participantID string) (float64, float64) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	participantPerformance := calculateParticipantPerformance(db.results, participantID)
	performances := collectPerformances(db.results)
	comparisonPercentage := calculateComparisonPercentage(participantPerformance, performances)

	return participantPerformance, comparisonPercentage
}

func calculateParticipantPerformance(results map[string]entities.Result, participantID string) float64 {
	participantResults := results[participantID]
	participantPerformance := 0.0
	if participantResults.TotalAnswers > 0 {
		participantPerformance = float64(participantResults.CorrectAnswers) / float64(participantResults.TotalAnswers) * 100
	}
	return participantPerformance
}

func collectPerformances(results map[string]entities.Result) []float64 {
	var performances []float64
	for _, result := range results {
		if result.TotalAnswers > 0 {
			performances = append(performances, float64(result.CorrectAnswers)/float64(result.TotalAnswers)*100)
		}
	}
	sort.Float64s(performances)
	return performances
}

func calculateComparisonPercentage(participantPerformance float64, performances []float64) float64 {
	betterCount := 0
	samePerformanceCount := 0
	for _, performance := range performances {
		if performance < participantPerformance {
			betterCount++
		} else if performance == participantPerformance {
			samePerformanceCount++
		}
	}

	comparisonPercentage := 0.0
	totalParticipants := len(performances)
	if totalParticipants > 0 {
		comparisonPercentage = (float64(betterCount) + float64(samePerformanceCount-1)/2.0) / float64(totalParticipants) * 100
	}

	return comparisonPercentage
}

func isAnswerCorrect(questions map[int]entities.Question, response entities.Response) bool {
	question, qExists := questions[response.QuestionID]
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
