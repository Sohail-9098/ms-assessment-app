package result

import (
	"testing"

	"github.com/Sohail-9098/ms-assessment-app/internal/assessment"
	"github.com/Sohail-9098/ms-assessment-app/internal/question"
	"github.com/stretchr/testify/require"
)

func TestResult_NewResultEmptyQuestionList(t *testing.T) {
	testAssessment := assessment.NewAssessment()
	expectedResult := Result{
		Assessment:    testAssessment,
		Correct:       0,
		Incorrect:     0,
		MarksTotal:    0,
		MarksObtained: 0,
	}
	actualResult := NewResult(testAssessment)
	require.Equal(t, expectedResult, actualResult)
}

func TestResult_NewResultWithQuestions(t *testing.T) {
	testAssessment := assessment.NewAssessment()
	testAssessment.Questions = question.QuestionSet
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
	q := question.QuestionSet[0]
	a := assessment.NewAssessment()
	a.Questions = append(a.Questions, q)
	r := NewResult(a)
	r.AssessAnswer(q, "Paris")
	require.Equal(t, 1, r.Correct)
	require.Equal(t, 0, r.Incorrect)
	require.Equal(t, 5, r.MarksObtained)
	require.Equal(t, 5, r.MarksTotal)
}

func TestResult_AccessAnswerInalid(t *testing.T) {
	q := question.QuestionSet[0]
	a := assessment.NewAssessment()
	a.Questions = append(a.Questions, q)
	r := NewResult(a)
	r.AssessAnswer(q, "London")
	require.Equal(t, 0, r.Correct)
	require.Equal(t, 1, r.Incorrect)
	require.Equal(t, 0, r.MarksObtained)
	require.Equal(t, 5, r.MarksTotal)
}
