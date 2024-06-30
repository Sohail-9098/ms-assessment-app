package question

type MultipleChoiceQuestion struct {
	QuestionId                   int
	QuestionText                 string
	Options                      []string
	PositiveMarks, NegativeMarks int
	correctAnswer                string
}

func (mcq *MultipleChoiceQuestion) IsCorrectAnswer(answer string) bool {
	return mcq.correctAnswer == answer
}

func (mcq *MultipleChoiceQuestion) CorrectAnswer() string {
	return mcq.correctAnswer
}

func (mcq *MultipleChoiceQuestion) SetCorrectAnswer(answer string) {
	mcq.correctAnswer = answer
}
