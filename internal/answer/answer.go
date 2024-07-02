package answer

type Answer struct {
	QuestionId int
	Answer     string
}

type UserAnswer struct {
	UserId       int
	AssessmentId int
	Answer       []Answer
}
