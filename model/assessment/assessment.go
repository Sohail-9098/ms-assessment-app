package assessment

import (
	"github.com/Sohail-9098/ms-assessment-app/model/question"
)

type Assessment struct {
	Questions []question.MultipleChoiceQuestion
	TimeLimit int
}
