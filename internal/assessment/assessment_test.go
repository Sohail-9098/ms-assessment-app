package assessment

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssessement_NewAssessment(t *testing.T) {
	expectedAssessment := Assessment{}
	actualAssessment := NewAssessment()
	require.Equal(t, expectedAssessment, actualAssessment)
}
