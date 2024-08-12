package get

type Result struct {
	CorrectAnswers int     `json:"correctAnswers"`
	TotalAnswers   int     `json:"totalAnswers"`
	Percentage     float64 `json:"percentage"`
}
