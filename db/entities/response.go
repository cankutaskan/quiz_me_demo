package entities

// Response represents the answer given by a participant
type Response struct {
	ParticipantID string // Unique identifier for the participant
	QuestionID    int    // ID of the question
	AnswerID      int    // ID of the selected answer
}
