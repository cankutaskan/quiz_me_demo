package entities

// Question represents a quiz question
type Question struct {
	ID             int
	Text           string
	Answers        []Answer
	AlternativeIDs []int // List of alternative question IDs
}
