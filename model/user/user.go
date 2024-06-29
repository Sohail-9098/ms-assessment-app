package user

import (
	"github.com/Sohail-9098/ms-assessment-app/model/result"
)

type User struct {
	UserId     int
	UserName   string
	TestResult []result.Result
}
