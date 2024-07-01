package result

import (
	"testing"

	"github.com/Sohail-9098/ms-assessment-app/internal/assessment"
	"github.com/Sohail-9098/ms-assessment-app/internal/question"
	"github.com/stretchr/testify/require"
)

var testQuestion question.MultipleChoiceQuestion = question.MultipleChoiceQuestion{
	QuestionId:    1,
	QuestionText:  "What is the capital of France",
	Options:       []string{"Delhi", "Stockholm", "Paris", "London"},
	PositiveMarks: 5,
	NegativeMarks: 0,
}

func TestResult_NewResultEmptyQuestionList(t *testing.T) {
	testAssessment := assessment.NewAssessment([]question.MultipleChoiceQuestion{testQuestion})
	expectedResult := Result{
		Assessment:    testAssessment,
		Correct:       0,
		Incorrect:     0,
		MarksTotal:    5,
		MarksObtained: 0,
	}
	actualResult := NewResult(testAssessment)
	require.Equal(t, expectedResult, actualResult)
}

func TestResult_NewResultWithQuestions(t *testing.T) {
	testAssessment := assessment.NewAssessment([]question.MultipleChoiceQuestion{testQuestion})
	testQuestion.SetCorrectAnswer("Paris")
	testAssessment.Questions = []question.MultipleChoiceQuestion{testQuestion}
	expectedResult := Result{
		Assessment:    testAssessment,
		Correct:       0,
		Incorrect:     0,
		MarksTotal:    len(testAssessment.Questions) * 5,
		MarksObtained: 0,
	}
	actualResult := NewResult(testAssessment)
	require.Equal(t, expectedResult, actualResult)
}

func TestResult_AccessAnswerValid(t *testing.T) {
	testQuestion.SetCorrectAnswer("Paris")
	q := testQuestion
	a := assessment.NewAssessment([]question.MultipleChoiceQuestion{testQuestion})
	r := NewResult(a)
	r.AssessAnswer(q, "Paris")
	require.Equal(t, 1, r.Correct)
	require.Equal(t, 0, r.Incorrect)
	require.Equal(t, 5, r.MarksObtained)
	require.Equal(t, 5, r.MarksTotal)
}

func TestResult_AccessAnswerInalid(t *testing.T) {
	testQuestion.SetCorrectAnswer("Paris")
	q := testQuestion
	a := assessment.NewAssessment([]question.MultipleChoiceQuestion{testQuestion})
	r := NewResult(a)
	r.AssessAnswer(q, "London")
	require.Equal(t, 0, r.Correct)
	require.Equal(t, 1, r.Incorrect)
	require.Equal(t, 0, r.MarksObtained)
	require.Equal(t, 5, r.MarksTotal)
}
