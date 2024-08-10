package entities

// Question represents a quiz question
type Question struct {
	ID       int
	Text     string
	Answers  []Answer
	Category string // Category for the question
}
