package db

import (
	"quiz_me/db/entities"
)

func Seed(db *DBContext) {
	// Add some questions
	questions := []entities.Question{
		{
			ID:   1,
			Text: "What is 2 + 2?",
			Answers: []entities.Answer{
				{ID: 1, Text: "3", IsCorrect: false},
				{ID: 2, Text: "4", IsCorrect: true},
				{ID: 3, Text: "5", IsCorrect: false},
				{ID: 4, Text: "6", IsCorrect: false},
			},
			AlternativeIDs: []int{2, 3}, // Alternatives in math
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
			AlternativeIDs: []int{1, 4}, // Alternatives in math
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
			AlternativeIDs: []int{1, 2}, // Alternatives in math
		},
		{
			ID:   4,
			Text: "What is the capital of Germany?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Berlin", IsCorrect: true},
				{ID: 2, Text: "Munich", IsCorrect: false},
				{ID: 3, Text: "Frankfurt", IsCorrect: false},
				{ID: 4, Text: "Hamburg", IsCorrect: false},
			},
			AlternativeIDs: []int{5, 6}, // Alternatives in geography
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
			AlternativeIDs: []int{4, 6}, // Alternatives in geography
		},
		{
			ID:   6,
			Text: "What is the capital of Italy?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Rome", IsCorrect: true},
				{ID: 2, Text: "Milan", IsCorrect: false},
				{ID: 3, Text: "Naples", IsCorrect: false},
				{ID: 4, Text: "Florence", IsCorrect: false},
			},
			AlternativeIDs: []int{4, 5}, // Alternatives in geography
		},
		{
			ID:   7,
			Text: "Who was the first President of the United States?",
			Answers: []entities.Answer{
				{ID: 1, Text: "George Washington", IsCorrect: true},
				{ID: 2, Text: "Thomas Jefferson", IsCorrect: false},
				{ID: 3, Text: "John Adams", IsCorrect: false},
				{ID: 4, Text: "James Madison", IsCorrect: false},
			},
			AlternativeIDs: []int{8, 9}, // Alternatives in history
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
			AlternativeIDs: []int{7, 9}, // Alternatives in history
		},
		{
			ID:   9,
			Text: "Who was the first Emperor of China?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Qin Shi Huang", IsCorrect: true},
				{ID: 2, Text: "Liu Bang", IsCorrect: false},
				{ID: 3, Text: "Wudi", IsCorrect: false},
				{ID: 4, Text: "Tang Taizong", IsCorrect: false},
			},
			AlternativeIDs: []int{7, 8}, // Alternatives in history
		},
		{
			ID:   10,
			Text: "What is the chemical symbol for water?",
			Answers: []entities.Answer{
				{ID: 1, Text: "H2O", IsCorrect: true},
				{ID: 2, Text: "O2", IsCorrect: false},
				{ID: 3, Text: "CO2", IsCorrect: false},
				{ID: 4, Text: "NaCl", IsCorrect: false},
			},
			AlternativeIDs: []int{11, 12}, // Alternatives in basic physics
		},
		{
			ID:   11,
			Text: "What is the chemical symbol for gold?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Au", IsCorrect: true},
				{ID: 2, Text: "Ag", IsCorrect: false},
				{ID: 3, Text: "Fe", IsCorrect: false},
				{ID: 4, Text: "Pb", IsCorrect: false},
			},
			AlternativeIDs: []int{10, 12}, // Alternatives in basic physics
		},
		{
			ID:   12,
			Text: "What is the chemical symbol for sodium?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Na", IsCorrect: true},
				{ID: 2, Text: "K", IsCorrect: false},
				{ID: 3, Text: "Ca", IsCorrect: false},
				{ID: 4, Text: "Mg", IsCorrect: false},
			},
			AlternativeIDs: []int{10, 11}, // Alternatives in basic physics
		},
		{
			ID:   13,
			Text: "What is the largest planet in our solar system?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Jupiter", IsCorrect: true},
				{ID: 2, Text: "Saturn", IsCorrect: false},
				{ID: 3, Text: "Earth", IsCorrect: false},
				{ID: 4, Text: "Mars", IsCorrect: false},
			},
			AlternativeIDs: []int{14, 15}, // Alternatives in basic physics
		},
		{
			ID:   14,
			Text: "What is the smallest planet in our solar system?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Mercury", IsCorrect: true},
				{ID: 2, Text: "Mars", IsCorrect: false},
				{ID: 3, Text: "Venus", IsCorrect: false},
				{ID: 4, Text: "Uranus", IsCorrect: false},
			},
			AlternativeIDs: []int{13, 15}, // Alternatives in basic physics
		},
		{
			ID:   15,
			Text: "What is the second largest planet in our solar system?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Saturn", IsCorrect: true},
				{ID: 2, Text: "Jupiter", IsCorrect: false},
				{ID: 3, Text: "Earth", IsCorrect: false},
				{ID: 4, Text: "Neptune", IsCorrect: false},
			},
			AlternativeIDs: []int{13, 14}, // Alternatives in basic physics
		},
		{
			ID:   16,
			Text: "What year did World War II begin?",
			Answers: []entities.Answer{
				{ID: 1, Text: "1939", IsCorrect: true},
				{ID: 2, Text: "1941", IsCorrect: false},
				{ID: 3, Text: "1945", IsCorrect: false},
				{ID: 4, Text: "1937", IsCorrect: false},
			},
			AlternativeIDs: []int{17, 18}, // Alternatives in history
		},
		{
			ID:   17,
			Text: "What year did the Titanic sink?",
			Answers: []entities.Answer{
				{ID: 1, Text: "1912", IsCorrect: true},
				{ID: 2, Text: "1905", IsCorrect: false},
				{ID: 3, Text: "1915", IsCorrect: false},
				{ID: 4, Text: "1920", IsCorrect: false},
			},
			AlternativeIDs: []int{16, 18}, // Alternatives in history
		},
		{
			ID:   18,
			Text: "What year did the Berlin Wall fall?",
			Answers: []entities.Answer{
				{ID: 1, Text: "1989", IsCorrect: true},
				{ID: 2, Text: "1991", IsCorrect: false},
				{ID: 3, Text: "1985", IsCorrect: false},
				{ID: 4, Text: "1979", IsCorrect: false},
			},
			AlternativeIDs: []int{16, 17}, // Alternatives in history
		},
		{
			ID:   19,
			Text: "What is the longest river in the world?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Nile", IsCorrect: true},
				{ID: 2, Text: "Amazon", IsCorrect: false},
				{ID: 3, Text: "Yangtze", IsCorrect: false},
				{ID: 4, Text: "Mississippi", IsCorrect: false},
			},
			AlternativeIDs: []int{20, 21}, // Alternatives in geography
		},
		{
			ID:   20,
			Text: "What is the largest desert in the world?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Sahara", IsCorrect: true},
				{ID: 2, Text: "Gobi", IsCorrect: false},
				{ID: 3, Text: "Kalahari", IsCorrect: false},
				{ID: 4, Text: "Arabian", IsCorrect: false},
			},
			AlternativeIDs: []int{19, 21}, // Alternatives in geography
		},
		{
			ID:   21,
			Text: "What is the largest ocean in the world?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Pacific Ocean", IsCorrect: true},
				{ID: 2, Text: "Atlantic Ocean", IsCorrect: false},
				{ID: 3, Text: "Indian Ocean", IsCorrect: false},
				{ID: 4, Text: "Arctic Ocean", IsCorrect: false},
			},
			AlternativeIDs: []int{19, 20}, // Alternatives in geography
		},
	}

	for _, q := range questions {
		db.AddQuestion(q)
	}
}
