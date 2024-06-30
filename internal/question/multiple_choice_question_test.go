package question

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultipleChoiceQuestion_CorrectAnswerValid(t *testing.T) {
	question := QuestionSet[0]
	actualAnswer := "Paris"
	require.Equal(t, true, question.IsCorrectAnswer(actualAnswer))
}

func TestMultipleChoiceQuestion_CorrectAnswerInvalid(t *testing.T) {
	question := QuestionSet[0]
	actualAnswer := "Lebanon"
	require.Equal(t, false, question.IsCorrectAnswer(actualAnswer))
}
