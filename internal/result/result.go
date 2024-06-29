package result

import (
	"github.com/Sohail-9098/ms-assessment-app/internal/assessment"
	"github.com/Sohail-9098/ms-assessment-app/internal/question"
)

type Result struct {
	Assessment                                    assessment.Assessment
	Correct, Incorrect, MarksTotal, MarksObtained int
}

func NewResult(assessment assessment.Assessment) Result {
	marksTotal := 0
	for _, question := range assessment.Questions {
		marksTotal += question.PositiveMarks
	}
	return Result{
		Assessment: assessment,
		MarksTotal: marksTotal,
	}
}

func (result *Result) AssessAnswer(question question.MultipleChoiceQuestion, answer string) {
	if question.IsCorrectAnswer(answer) {
		result.Correct++
		result.MarksObtained += question.PositiveMarks
	} else {
		result.Incorrect++
		result.MarksObtained -= question.NegativeMarks
	}
}
