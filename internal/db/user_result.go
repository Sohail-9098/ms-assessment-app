package db

import (
	"log"

	"github.com/Sohail-9098/ms-assessment-app/internal/result"
	"github.com/Sohail-9098/ms-assessment-app/internal/user"
)

func AddUserResult(user user.User, testResult result.Result) {
	conn := connect()
	defer close(conn)

	query := "INSERT INTO user_results (user_id, assessment_id, correct_answers, incorrect_answers, marks_total, marks_obtained) VALUES ($1,$2,$3,$4,$5,$6);"

	_, err := conn.Exec(query, user.UserId, testResult.Assessment.AssessmentId, testResult.Correct, testResult.Incorrect, testResult.MarksTotal, testResult.MarksObtained)
	if err != nil {
		log.Fatalf("error updating result to db: %v", err)
	}
}

func GetUserResults(user *user.User) []result.Result {
	conn := connect()
	defer close(conn)

	query := "SELECT assessment_id, correct_answers, incorrect_answers, marks_total, marks_obtained FROM user_results WHERE user_id=$1;"

	rows, err := conn.Query(query, user.UserId)
	if err != nil {
		log.Fatalf("error fetching question from db: %v", err)
	}

	userResults := []result.Result{}
	var assessmentId, correct, incorrect, total, obtained int32

	for rows.Next() {
		err := rows.Scan(&assessmentId, &correct, &incorrect, &total, &obtained)
		if err != nil {
			log.Fatalf("error reading result set: %v", err)
		}
		userResult := result.Result{
			Assessment:    GetAssessmentById(int(assessmentId)),
			Correct:       int(correct),
			Incorrect:     int(incorrect),
			MarksTotal:    int(total),
			MarksObtained: int(obtained),
		}
		userResults = append(userResults, userResult)
	}

	user.TestResult = userResults

	return userResults
}
