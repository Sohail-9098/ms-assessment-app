package main

import (
	"fmt"

	"github.com/Sohail-9098/ms-assessment-app/model/assessment"
	"github.com/Sohail-9098/ms-assessment-app/model/question"
	"github.com/Sohail-9098/ms-assessment-app/model/result"
	"github.com/Sohail-9098/ms-assessment-app/model/user"
)

func main() {
	fmt.Println("Welcome to Assessment App!")
	newUser := user.User{UserId: 1, UserName: "Tarannum Ali"}
	newAssessment := assessment.Assessment{Questions: question.QuestionSet, TimeLimit: 30}
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
			userResult.MarksObtained += question.Mark
		} else {
			userResult.MarksObtained -= 1
		}
	}
	newUser.TestResult = append(newUser.TestResult, userResult)
	fmt.Println("Your Result: ")
	fmt.Printf("User Name: %v, User Score: %v\n", newUser.UserName, newUser.TestResult[0].MarksObtained)
	fmt.Printf("%d/%d\n", userResult.MarksObtained, userResult.MarksTotal)
	fmt.Printf("%v percent \n", userResult.MarksObtained*100/userResult.MarksTotal)
}
