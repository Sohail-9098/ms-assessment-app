package db

import (
	"log"

	"github.com/Sohail-9098/ms-assessment-app/internal/assessment"
	"github.com/Sohail-9098/ms-assessment-app/internal/question"
	"github.com/Sohail-9098/ms-assessment-app/internal/user"
	"github.com/lib/pq"
)

func AddAssessment(newAssessment assessment.Assessment) {
	conn := connect()
	defer close(conn)

	query := "INSERT INTO assessments (questions) VALUES ($1);"

	questionIds := []int{}

	for _, ques := range newAssessment.Questions {
		questionIds = append(questionIds, ques.QuestionId)
	}

	_, err := conn.Exec(query, pq.Array(questionIds))
	if err != nil {
		log.Fatalf("error assessment to db: %v", err)
	}

}

func GetAssessmentById(assessmentId int) assessment.Assessment {
	conn := connect()
	defer close(conn)

	query := "SELECT * FROM assessments WHERE assessment_id=$1;"

	rows, err := conn.Query(query, assessmentId)
	if err != nil {
		log.Fatalf("error fetching assessment from db: %v", err)
	}

	newAssessment := assessment.Assessment{}
	var rawIds []int32
	for rows.Next() {
		err := rows.Scan(&newAssessment.AssessmentId, pq.Array(&rawIds))
		if err != nil {
			log.Fatalf("error reading result set: %v", err)
		}
	}

	var questionIds []int

	for _, id := range rawIds {
		questionIds = append(questionIds, int(id))
	}

	questions := []question.MultipleChoiceQuestion{}
	for _, id := range questionIds {
		questions = append(questions, GetMCQById(id))
	}

	newAssessment.Questions = questions
	return newAssessment
}

func AssignAssessment(user user.User, assessment assessment.Assessment) {
	conn := connect()
	defer close(conn)

	query := "INSERT INTO user_assignment VALUES($1, $2);"
	_, err := conn.Exec(query, user.UserId, assessment.AssessmentId)
	if err != nil {
		log.Fatalf("error inserting assignment : %v", err)
	}

}
