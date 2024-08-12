package entities

type Question struct {
	ID       int
	Text     string
	Answers  []Answer
	Category string
}
