package result

import "github.com/Sohail-9098/ms-assessment-app/model/assessment"

type Result struct {
	Assessment    assessment.Assessment
	MarksTotal    int
	MarksObtained int
}
