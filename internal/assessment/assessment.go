package assessment

import (
	"github.com/Sohail-9098/ms-assessment-app/internal/question"
)

var index int = 0

type Assessment struct {
	AssessmentId int
	Questions    []question.MultipleChoiceQuestion
}

func NewAssessment(questions []question.MultipleChoiceQuestion) Assessment {
	index++
	return Assessment{
		AssessmentId: index,
		Questions:    questions,
	}
}
