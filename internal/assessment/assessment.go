package assessment

import (
	"github.com/Sohail-9098/ms-assessment-app/internal/question"
)

type Assessment struct {
	AssessmentId int
	Questions    []question.MultipleChoiceQuestion
}

func NewAssessment(questions []question.MultipleChoiceQuestion) Assessment {
	return Assessment{
		Questions: questions,
	}
}
