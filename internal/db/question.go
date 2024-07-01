package db

import (
	"log"

	"github.com/Sohail-9098/ms-assessment-app/internal/question"
	"github.com/lib/pq"
)

func AddMCQ(mcq question.MultipleChoiceQuestion) {
	conn := connect()
	defer close(conn)

	query := "INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ($1,$2,$3,$4,$5);"

	_, err := conn.Exec(query, mcq.QuestionText, pq.Array(mcq.Options), mcq.PositiveMarks, mcq.NegativeMarks, mcq.CorrectAnswer())
	if err != nil {
		log.Fatalf("error adding question to db: %v", err)
	}

	log.Println("question added")
}

func GetMCQById(questionId int) question.MultipleChoiceQuestion {
	conn := connect()
	defer close(conn)

	query := "SELECT * FROM multiple_choice_questions WHERE question_id=$1;"

	rows, err := conn.Query(query, questionId)
	if err != nil {
		log.Fatalf("error fetching question from db: %v", err)
	}

	mcq := question.MultipleChoiceQuestion{}
	correctAnswer := ""
	for rows.Next() {
		err := rows.Scan(&mcq.QuestionId, &mcq.QuestionText, pq.Array(&mcq.Options), &mcq.PositiveMarks, &mcq.NegativeMarks, &correctAnswer)
		if err != nil {
			log.Fatalf("error reading result set: %v", err)
		}
	}

	mcq.SetCorrectAnswer(correctAnswer)
	return mcq
}

func GetMCQs() []question.MultipleChoiceQuestion {
	conn := connect()
	defer close(conn)

	query := "SELECT * FROM multiple_choice_questions;"

	rows, err := conn.Query(query)
	if err != nil {
		log.Fatalf("error fetching question from db: %v", err)
	}

	mcq := question.MultipleChoiceQuestion{}
	mcqs := []question.MultipleChoiceQuestion{}
	correctAnswer := ""
	for rows.Next() {
		err := rows.Scan(&mcq.QuestionId, &mcq.QuestionText, pq.Array(&mcq.Options), &mcq.PositiveMarks, &mcq.NegativeMarks, &correctAnswer)
		if err != nil {
			log.Fatalf("error reading result set: %v", err)
		}
		mcq.SetCorrectAnswer(correctAnswer)
		mcqs = append(mcqs, mcq)
	}
	return mcqs
}
