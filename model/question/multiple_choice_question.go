package question

type MultipleChoiceQuestion struct {
	
	Options       [4]string
	Mark          int
	correctAnswer string
}

func (mcq *MultipleChoiceQuestion) CorrectAnswer() string {
	return mcq.correctAnswer
}
