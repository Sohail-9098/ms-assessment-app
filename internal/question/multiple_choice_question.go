package question

type MultipleChoiceQuestion struct {
	QuestionId                   int
	QuestionText                 string
	Options                      []string
	PositiveMarks, NegativeMarks int
	correctAnswer                string
}

func NewQuestion(questionText string, options []string, positiveMarks, negativeMarks int, correctAnswer string) MultipleChoiceQuestion {
	return MultipleChoiceQuestion{
		QuestionText:  questionText,
		Options:       options,
		PositiveMarks: positiveMarks, NegativeMarks: negativeMarks,
		correctAnswer: correctAnswer,
	}
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

func (mcq *MultipleChoiceQuestion) AddQuestionToDb() {
	mcq.addQuestionToDb()
}
