package question

type MultipleChoiceQuestion struct {
	QuestionText  string
	Options       []string
	Mark          int
	correctAnswer string
}

func (mcq *MultipleChoiceQuestion) CorrectAnswer() string {
	return mcq.correctAnswer
}
