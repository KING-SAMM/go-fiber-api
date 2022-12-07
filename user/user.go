package user

import (
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type User struct {
	gorm.Model 
	FirstName string	`json:"firstname"`
	LastName string		`json:"lastname"`
	Email string		`json:"email"`
}