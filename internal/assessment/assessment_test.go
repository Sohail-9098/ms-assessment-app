package assessment

import (
	"testing"

	"github.com/Sohail-9098/ms-assessment-app/internal/question"
	"github.com/stretchr/testify/require"
)

func TestAssessement_NewAssessment(t *testing.T) {
	expectedAssessment1 := Assessment{0, testQuestionSet}
	expectedAssessment2 := Assessment{0, testQuestionSet}
	actualAssessment1 := NewAssessment(testQuestionSet)
	actualAssessment2 := NewAssessment(testQuestionSet)
	require.Equal(t, expectedAssessment1, actualAssessment1)
	require.Equal(t, expectedAssessment2, actualAssessment2)
}

var testQuestionSet = []question.MultipleChoiceQuestion{
	{
		QuestionText:  "What is the capital of France?",
		Options:       []string{"Berlin", "Madrid", "Paris", "Lisbon"},
		PositiveMarks: 5,
		NegativeMarks: 0,
	},
}
