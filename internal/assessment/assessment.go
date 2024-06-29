package assessment

import (
	"github.com/Sohail-9098/ms-assessment-app/internal/question"
)

type Assessment struct {
	Questions []question.MultipleChoiceQuestion
}

func NewAssessment() Assessment {
	return Assessment{}
}
