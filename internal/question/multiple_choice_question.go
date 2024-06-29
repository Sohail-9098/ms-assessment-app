package question

type MultipleChoiceQuestion struct {
	QuestionText                 string
	Options                      []string
	PositiveMarks, NegativeMarks int
	correctAnswer                string
}

func (mcq *MultipleChoiceQuestion) IsCorrectAnswer(answer string) bool {
	return mcq.correctAnswer == answer
}
