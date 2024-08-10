package db

import (
	"quiz_me/db/entities"
)

func Seed(db *DBContext) {
	// Add some questions
	questions := []entities.Question{
		// Math Questions
		{
			ID:   1,
			Text: "What is 2 + 2?",
			Answers: []entities.Answer{
				{ID: 1, Text: "3", IsCorrect: false},
				{ID: 2, Text: "4", IsCorrect: true},
				{ID: 3, Text: "5", IsCorrect: false},
				{ID: 4, Text: "6", IsCorrect: false},
			},
			Category: "math",
		},
		{
			ID:   2,
			Text: "What is 3 + 5?",
			Answers: []entities.Answer{
				{ID: 1, Text: "7", IsCorrect: false},
				{ID: 2, Text: "8", IsCorrect: true},
				{ID: 3, Text: "9", IsCorrect: false},
				{ID: 4, Text: "10", IsCorrect: false},
			},
			Category: "math",
		},
		{
			ID:   3,
			Text: "What is 10 - 4?",
			Answers: []entities.Answer{
				{ID: 1, Text: "5", IsCorrect: false},
				{ID: 2, Text: "6", IsCorrect: true},
				{ID: 3, Text: "7", IsCorrect: false},
				{ID: 4, Text: "8", IsCorrect: false},
			},
			Category: "math",
		},

		// Geography Questions
		{
			ID:   4,
			Text: "What is the capital of Germany?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Berlin", IsCorrect: true},
				{ID: 2, Text: "Munich", IsCorrect: false},
				{ID: 3, Text: "Frankfurt", IsCorrect: false},
				{ID: 4, Text: "Hamburg", IsCorrect: false},
			},
			Category: "geography",
		},
		{
			ID:   5,
			Text: "What is the capital of Spain?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Madrid", IsCorrect: true},
				{ID: 2, Text: "Barcelona", IsCorrect: false},
				{ID: 3, Text: "Seville", IsCorrect: false},
				{ID: 4, Text: "Valencia", IsCorrect: false},
			},
			Category: "geography",
		},
		{
			ID:   6,
			Text: "What is the largest desert in the world?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Sahara", IsCorrect: true},
				{ID: 2, Text: "Gobi", IsCorrect: false},
				{ID: 3, Text: "Kalahari", IsCorrect: false},
				{ID: 4, Text: "Arabian", IsCorrect: false},
			},
			Category: "geography",
		},

		// History Questions
		{
			ID:   7,
			Text: "Who was the first President of the United States?",
			Answers: []entities.Answer{
				{ID: 1, Text: "George Washington", IsCorrect: true},
				{ID: 2, Text: "Thomas Jefferson", IsCorrect: false},
				{ID: 3, Text: "John Adams", IsCorrect: false},
				{ID: 4, Text: "James Madison", IsCorrect: false},
			},
			Category: "history",
		},
		{
			ID:   8,
			Text: "Who was the first man to walk on the moon?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Neil Armstrong", IsCorrect: true},
				{ID: 2, Text: "Buzz Aldrin", IsCorrect: false},
				{ID: 3, Text: "Yuri Gagarin", IsCorrect: false},
				{ID: 4, Text: "Michael Collins", IsCorrect: false},
			},
			Category: "history",
		},
		{
			ID:   9,
			Text: "What year did World War II begin?",
			Answers: []entities.Answer{
				{ID: 1, Text: "1939", IsCorrect: true},
				{ID: 2, Text: "1941", IsCorrect: false},
				{ID: 3, Text: "1945", IsCorrect: false},
				{ID: 4, Text: "1937", IsCorrect: false},
			},
			Category: "history",
		},
	}

	for _, q := range questions {
		db.AddQuestion(q)
	}
}
