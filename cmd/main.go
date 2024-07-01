package main

import (
	"fmt"

	"github.com/Sohail-9098/ms-assessment-app/internal/db"
	"github.com/Sohail-9098/ms-assessment-app/internal/result"
)

func main() {
	fmt.Println("Welcome to Assessment App!")
	// User LogsIn
	newUser := db.GetUserId(1)

	// Creates Assessment and save it to db
	// newAssessment := assessment.NewAssessment([]question.MultipleChoiceQuestion{db.GetMCQById(1), db.GetMCQById(2), db.GetMCQById(3)})
	// db.AddAssessment(newAssessment)

	// Fetch Assessment
	assessmentToTake := db.GetAssessmentById(1)

	// Create User Result
	assessmentResult := result.NewResult(assessmentToTake)

	// Start the assessment
	for _, question := range assessmentToTake.Questions {
		// Ask Question
		fmt.Println(question.QuestionText)

		for i, option := range question.Options {
			fmt.Printf("%d: %v\n", i+1, option)
		}
		// Get user answer
		var answerIndex int
		fmt.Scanln(&answerIndex)
		selectedAnswer := question.Options[answerIndex-1]
		assessmentResult.AssessAnswer(question, selectedAnswer)
	}

	// update the db with results
	db.AddUserResult(newUser, assessmentResult)

	db.GetUserResults(&newUser)

	// Print results
	fmt.Println("Your Results: ")
	for _, result := range newUser.TestResult {
		fmt.Printf("User Name: %v, Correct: %v, Incorrect: %v\n", newUser.UserName, result.Correct, result.Incorrect)
		fmt.Printf("Total Score: %d/%d\n", result.MarksObtained, result.MarksTotal)
	}
}
