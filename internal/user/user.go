package user

import (
	"github.com/Sohail-9098/ms-assessment-app/internal/result"
)

var index int = 0

type User struct {
	UserId     int
	UserName   string
	TestResult []result.Result
}

func NewUser() User {
	index++
	return User{
		UserId: index,
	}
}
