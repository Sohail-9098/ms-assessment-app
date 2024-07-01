package question

import (
	"log"

	"github.com/Sohail-9098/ms-assessment-app/internal/db"
	"github.com/lib/pq"
)

func (mcq *MultipleChoiceQuestion) addQuestionToDb() {
	conn := db.Connect()
	defer db.Close(conn)

	query := "INSERT INTO multiple_choice_questions (question_text, options, positive_marks, negative_marks, correct_answer) VALUES ($1,$2,$3,$4,$5);"

	_, err := conn.Exec(query, mcq.QuestionText, pq.Array(mcq.Options), mcq.PositiveMarks, mcq.NegativeMarks, mcq.correctAnswer)
	if err != nil {
		log.Fatalf("error adding question to db: %v", err)
	}

	log.Println("question added")
}
