package db

import (
	"log"
)

func Seed(db *InMemoryDB) {
	// Add some questions
	q1 := Question{
		ID:   1,
		Text: "What is 2 + 2?",
		Answers: []Answer{
			{Text: "3", IsCorrect: false},
			{Text: "4", IsCorrect: true},
			{Text: "5", IsCorrect: false},
			{Text: "6", IsCorrect: false},
		},
		AlternativeIDs: []int{},
	}

	q2 := Question{
		ID:   2,
		Text: "What is the capital of France?",
		Answers: []Answer{
			{Text: "London", IsCorrect: false},
			{Text: "Paris", IsCorrect: true},
			{Text: "Berlin", IsCorrect: false},
			{Text: "Madrid", IsCorrect: false},
		},
		AlternativeIDs: []int{},
	}

	db.AddQuestion(q1)
	db.AddQuestion(q2)

	// Add some responses
	db.AddResponse(Response{ParticipantID: "user1", QuestionID: 1, SelectedAnswerIndex: 1})
	db.AddResponse(Response{ParticipantID: "user1", QuestionID: 2, SelectedAnswerIndex: 1})

	db.AddResponse(Response{ParticipantID: "user2", QuestionID: 1, SelectedAnswerIndex: 0})
	db.AddResponse(Response{ParticipantID: "user2", QuestionID: 2, SelectedAnswerIndex: 1})

	log.Println("Database seeded with initial questions and responses")
}
