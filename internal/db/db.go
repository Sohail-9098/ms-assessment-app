package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Sohail-9098/ms-assessment-app/internal/question"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type dbConfig struct {
	User, Password, DbName, SslMode string
}

func main() {
	config := dbConfig{"postgres", "pg@123", "ms_assessment_app", "disable"}
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", config.User, config.Password, config.DbName, config.SslMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer db.Close()

	fmt.Println("Working")

	query := "SELECT * FROM multiple_choice_questions;"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("failed to execute query: %v", err)
	}

	mcq := question.MultipleChoiceQuestion{}
	mcqs := []question.MultipleChoiceQuestion{}
	var correctAnswer string
	for rows.Next() {
		err := rows.Scan(&mcq.QuestionId, &mcq.QuestionText, pq.Array(&mcq.Options), &mcq.PositiveMarks, &mcq.NegativeMarks, &correctAnswer)
		if err != nil {
			log.Fatalf("failed to iterate: %v", err)
		}
		mcq.SetCorrectAnswer(correctAnswer)
		mcqs = append(mcqs, mcq)
	}

	for _, mcq := range mcqs{
		fmt.Println(mcq)
	}
}
