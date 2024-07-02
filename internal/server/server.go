package server

import (
	"net/http"
	"strconv"

	"github.com/Sohail-9098/ms-assessment-app/internal/db"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()
	r.GET("/assignments/:userId", getAssignments)
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
