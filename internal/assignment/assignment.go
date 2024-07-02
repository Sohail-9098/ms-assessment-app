package assignment

import (
	"github.com/Sohail-9098/ms-assessment-app/internal/assessment"
	"github.com/Sohail-9098/ms-assessment-app/internal/db"
	"github.com/Sohail-9098/ms-assessment-app/internal/user"
)

type Assignment struct {
	AssignmentId int
	User         user.User
	Assessment   assessment.Assessment
}

func (a Assignment) AssignAssessment() {
	db.AssignAssessment(a.User, a.Assessment)
}
