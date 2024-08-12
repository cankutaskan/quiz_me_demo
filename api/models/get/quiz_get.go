package get

import "quiz_me/db/entities"

// API models
type Answer struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type Question struct {
	ID      int      `json:"id"`
	Text    string   `json:"text"`
	Answers []Answer `json:"answers"`
}

type Quiz struct {
	Questions []Question `json:"questions"`
}

func Convert(questions []entities.Question) Quiz {
	var apiQuestions []Question
	for _, q := range questions {
		var apiAnswers []Answer
		for _, a := range q.Answers {
			apiAnswers = append(apiAnswers, Answer{
				ID:   a.ID,
				Text: a.Text,
			})
		}
		apiQuestions = append(apiQuestions, Question{
			ID:      q.ID,
			Text:    q.Text,
			Answers: apiAnswers,
		})
	}
	return Quiz{
		Questions: apiQuestions,
	}
}
