package main

import (
	"fmt"

	"github.com/Sohail-9098/ms-assessment-app/internal/assessment"
	"github.com/Sohail-9098/ms-assessment-app/internal/db"
	"github.com/Sohail-9098/ms-assessment-app/internal/question"
	"github.com/Sohail-9098/ms-assessment-app/internal/result"
	"github.com/Sohail-9098/ms-assessment-app/internal/user"
)

func main() {
	fmt.Println("Welcome to Assessment App!")
	// User LogsIn
	newUser := user.NewUser()
	newUser.UserId = 1
	newUser.UserName = "John Wick"

	// Creates Assessment and save it to db
	newAssessment := assessment.NewAssessment([]question.MultipleChoiceQuestion{db.GetMCQById(1), db.GetMCQById(2), db.GetMCQById(3)})
	db.AddAssessment(newAssessment)

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

	// Add Result to user profile
	newUser.TestResult = append(newUser.TestResult, assessmentResult)

	// Print results
	fmt.Println("Your Result: ")
	fmt.Printf("User Name: %v, Correct: %v, Incorrect: %v\n", newUser.UserName, newUser.TestResult[0].Correct, newUser.TestResult[0].Incorrect)
	fmt.Printf("Total Score: %d/%d\n", assessmentResult.MarksObtained, assessmentResult.MarksTotal)
}
