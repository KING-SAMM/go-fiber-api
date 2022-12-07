package user

import (
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:admin@tcp(127.0.0.1:3306)/godb?charset=utf8"

type User struct {
	gorm.Model 
	FirstName string	`json:"firstname"`
	LastName string		`json:"lastname"`
	Email string		`json:"email"`
}