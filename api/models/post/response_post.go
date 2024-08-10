package post

type AnswerResponse struct {
	QuestionID int `json:"question_id"`
	AnswerID   int `json:"answer_id"`
}

type Responses struct {
	UserID    string           `json:"user_id"`
	Responses []AnswerResponse `json:"responses"`
}
