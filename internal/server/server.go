package server

import (
	"net/http"
	"strconv"

	"github.com/Sohail-9098/ms-assessment-app/internal/answer"
	"github.com/Sohail-9098/ms-assessment-app/internal/db"
	"github.com/Sohail-9098/ms-assessment-app/internal/result"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()
	r.GET("/assignments/:userId", getAssignments)
	r.POST("/submit", submitTest)
	r.Run(":8080")
}

func getAssignments(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user := db.GetUserById(userId)
	assessment := db.GetAssignment(user)
	c.JSON(http.StatusOK, assessment)
}

func submitTest(c *gin.Context) {
	var userAnswers answer.UserAnswer
	if err := c.ShouldBindJSON(&userAnswers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := calculateResult(userAnswers)
	c.JSON(http.StatusOK, results)
}

func calculateResult(userAnswers answer.UserAnswer) result.Result {
	assessment := db.GetAssessmentById(userAnswers.AssessmentId)

	results := result.NewResult(assessment)

	for _, answer := range userAnswers.Answer {
		question := db.GetMCQById(answer.QuestionId)
		if question.IsCorrectAnswer(answer.Answer) {
			results.Correct++
			results.MarksObtained += question.PositiveMarks
		} else {
			results.Incorrect++
			results.MarksObtained -= question.NegativeMarks
		}
	}
	user := db.GetUserById(userAnswers.UserId)
	db.AddUserResult(user, results)

	return results
}
