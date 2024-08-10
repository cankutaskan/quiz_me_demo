package db

import (
	"quiz_me/db/entities"
)

func Seed(db *DBContext) {
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
		{
			ID:   10,
			Text: "What is 15 x 3?",
			Answers: []entities.Answer{
				{ID: 1, Text: "30", IsCorrect: false},
				{ID: 2, Text: "45", IsCorrect: true},
				{ID: 3, Text: "50", IsCorrect: false},
				{ID: 4, Text: "60", IsCorrect: false},
			},
			Category: "math",
		},
		{
			ID:   11,
			Text: "What is the square root of 81?",
			Answers: []entities.Answer{
				{ID: 1, Text: "7", IsCorrect: false},
				{ID: 2, Text: "8", IsCorrect: false},
				{ID: 3, Text: "9", IsCorrect: true},
				{ID: 4, Text: "10", IsCorrect: false},
			},
			Category: "math",
		},
		{
			ID:   12,
			Text: "What is 7 + 6?",
			Answers: []entities.Answer{
				{ID: 1, Text: "12", IsCorrect: false},
				{ID: 2, Text: "13", IsCorrect: true},
				{ID: 3, Text: "14", IsCorrect: false},
				{ID: 4, Text: "15", IsCorrect: false},
			},
			Category: "math",
		},
		{
			ID:   13,
			Text: "What is 25 - 7?",
			Answers: []entities.Answer{
				{ID: 1, Text: "17", IsCorrect: false},
				{ID: 2, Text: "18", IsCorrect: true},
				{ID: 3, Text: "19", IsCorrect: false},
				{ID: 4, Text: "20", IsCorrect: false},
			},
			Category: "math",
		},
		{
			ID:   14,
			Text: "What is 9 / 3?",
			Answers: []entities.Answer{
				{ID: 1, Text: "2", IsCorrect: false},
				{ID: 2, Text: "3", IsCorrect: true},
				{ID: 3, Text: "4", IsCorrect: false},
				{ID: 4, Text: "5", IsCorrect: false},
			},
			Category: "math",
		},
		{
			ID:   15,
			Text: "What is 4^2?",
			Answers: []entities.Answer{
				{ID: 1, Text: "8", IsCorrect: false},
				{ID: 2, Text: "12", IsCorrect: false},
				{ID: 3, Text: "16", IsCorrect: true},
				{ID: 4, Text: "20", IsCorrect: false},
			},
			Category: "math",
		},
		{
			ID:   16,
			Text: "What is the cube root of 27?",
			Answers: []entities.Answer{
				{ID: 1, Text: "2", IsCorrect: false},
				{ID: 2, Text: "3", IsCorrect: true},
				{ID: 3, Text: "4", IsCorrect: false},
				{ID: 4, Text: "5", IsCorrect: false},
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
		{
			ID:   17,
			Text: "Which is the largest continent by area?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Africa", IsCorrect: false},
				{ID: 2, Text: "Asia", IsCorrect: true},
				{ID: 3, Text: "Europe", IsCorrect: false},
				{ID: 4, Text: "South America", IsCorrect: false},
			},
			Category: "geography",
		},
		{
			ID:   18,
			Text: "What is the smallest country in the world?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Monaco", IsCorrect: false},
				{ID: 2, Text: "Vatican City", IsCorrect: true},
				{ID: 3, Text: "San Marino", IsCorrect: false},
				{ID: 4, Text: "Liechtenstein", IsCorrect: false},
			},
			Category: "geography",
		},
		{
			ID:   19,
			Text: "Which river is the longest in the world?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Amazon", IsCorrect: false},
				{ID: 2, Text: "Nile", IsCorrect: true},
				{ID: 3, Text: "Yangtze", IsCorrect: false},
				{ID: 4, Text: "Mississippi", IsCorrect: false},
			},
			Category: "geography",
		},
		{
			ID:   20,
			Text: "What is the capital of Australia?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Sydney", IsCorrect: false},
				{ID: 2, Text: "Melbourne", IsCorrect: false},
				{ID: 3, Text: "Canberra", IsCorrect: true},
				{ID: 4, Text: "Brisbane", IsCorrect: false},
			},
			Category: "geography",
		},
		{
			ID:   21,
			Text: "Which ocean is the largest?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Atlantic", IsCorrect: false},
				{ID: 2, Text: "Pacific", IsCorrect: true},
				{ID: 3, Text: "Indian", IsCorrect: false},
				{ID: 4, Text: "Southern", IsCorrect: false},
			},
			Category: "geography",
		},
		{
			ID:   22,
			Text: "Which country has the most population?",
			Answers: []entities.Answer{
				{ID: 1, Text: "India", IsCorrect: false},
				{ID: 2, Text: "China", IsCorrect: true},
				{ID: 3, Text: "United States", IsCorrect: false},
				{ID: 4, Text: "Indonesia", IsCorrect: false},
			},
			Category: "geography",
		},
		{
			ID:   23,
			Text: "What is the capital of Canada?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Toronto", IsCorrect: false},
				{ID: 2, Text: "Ottawa", IsCorrect: true},
				{ID: 3, Text: "Vancouver", IsCorrect: false},
				{ID: 4, Text: "Montreal", IsCorrect: false},
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
		{
			ID:   24,
			Text: "Who wrote the Declaration of Independence?",
			Answers: []entities.Answer{
				{ID: 1, Text: "George Washington", IsCorrect: false},
				{ID: 2, Text: "Thomas Jefferson", IsCorrect: true},
				{ID: 3, Text: "John Adams", IsCorrect: false},
				{ID: 4, Text: "Benjamin Franklin", IsCorrect: false},
			},
			Category: "history",
		},
		{
			ID:   25,
			Text: "In which year did the Titanic sink?",
			Answers: []entities.Answer{
				{ID: 1, Text: "1905", IsCorrect: false},
				{ID: 2, Text: "1912", IsCorrect: true},
				{ID: 3, Text: "1918", IsCorrect: false},
				{ID: 4, Text: "1921", IsCorrect: false},
			},
			Category: "history",
		},
		{
			ID:   26,
			Text: "Who discovered America?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Christopher Columbus", IsCorrect: true},
				{ID: 2, Text: "Vasco da Gama", IsCorrect: false},
				{ID: 3, Text: "Marco Polo", IsCorrect: false},
				{ID: 4, Text: "Leif Erikson", IsCorrect: false},
			},
			Category: "history",
		},
		{
			ID:   27,
			Text: "Which year was the Berlin Wall torn down?",
			Answers: []entities.Answer{
				{ID: 1, Text: "1987", IsCorrect: false},
				{ID: 2, Text: "1989", IsCorrect: true},
				{ID: 3, Text: "1991", IsCorrect: false},
				{ID: 4, Text: "1993", IsCorrect: false},
			},
			Category: "history",
		},
		{
			ID:   28,
			Text: "Who was the first Emperor of Rome?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Julius Caesar", IsCorrect: false},
				{ID: 2, Text: "Augustus", IsCorrect: true},
				{ID: 3, Text: "Nero", IsCorrect: false},
				{ID: 4, Text: "Tiberius", IsCorrect: false},
			},
			Category: "history",
		},
		{
			ID:   29,
			Text: "What was the name of the ship on which the Pilgrims traveled to America?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Mayflower", IsCorrect: true},
				{ID: 2, Text: "Santa Maria", IsCorrect: false},
				{ID: 3, Text: "Victory", IsCorrect: false},
				{ID: 4, Text: "Beagle", IsCorrect: false},
			},
			Category: "history",
		},
		{
			ID:   30,
			Text: "Who was the British Prime Minister during World War II?",
			Answers: []entities.Answer{
				{ID: 1, Text: "Winston Churchill", IsCorrect: true},
				{ID: 2, Text: "Neville Chamberlain", IsCorrect: false},
				{ID: 3, Text: "Clement Attlee", IsCorrect: false},
				{ID: 4, Text: "Stanley Baldwin", IsCorrect: false},
			},
			Category: "history",
		},
	}

	for _, q := range questions {
		db.AddQuestion(q)
	}
}
