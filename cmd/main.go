package main

import (
	"fmt"

	"github.com/Sohail-9098/ms-assessment-app/internal/assessment"
	"github.com/Sohail-9098/ms-assessment-app/internal/question"
	"github.com/Sohail-9098/ms-assessment-app/internal/result"
	"github.com/Sohail-9098/ms-assessment-app/internal/user"
)

func main() {
	fmt.Println("Welcome to Assessment App!")
	newUser := user.User{UserId: 1, UserName: "Tarannum Ali"}
	newAssessment := assessment.Assessment{Questions: question.QuestionSet[:2]}
	userResult := result.Result{Assessment: newAssessment}
	for _, question := range newAssessment.Questions {
		fmt.Println(question.QuestionText)
		userResult.MarksTotal += question.Mark
		for i, option := range question.Options {
			fmt.Printf("%d: %v\n", i+1, option)
		}
		var answer int
		fmt.Scanln(&answer)
		if question.Options[answer-1] == question.CorrectAnswer() {
			userResult.Correct++
			userResult.MarksObtained += question.Mark
		} else {
			userResult.Incorrect++
			userResult.MarksObtained -= 1
		}
	}
	newUser.TestResult = append(newUser.TestResult, userResult)
	fmt.Println("Your Result: ")
	fmt.Printf("User Name: %v, Correct: %v, Incorrect: %v\n", newUser.UserName, newUser.TestResult[0].Correct, newUser.TestResult[0].Incorrect)
	fmt.Printf("Total Score: %d/%d\n", userResult.MarksObtained, userResult.MarksTotal)
}
