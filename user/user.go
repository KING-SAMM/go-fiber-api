package user

import (
	"fmt"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:admin@tcp(127.0.0.1:3308)/godb?charset=utf8"

type User struct {
	gorm.Model 
	FirstName string	`json:"firstname"`
	LastName string		`json:"lastname"`
	Email string		`json:"email"`
}

func initialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config())

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database!")
	}
	DB.AutoMigrate(&User{})
}

func GetUsers(c *fiber.Ctx) error {

}
func GetUser(c *fiber.Ctx) error {

}
func SaveUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&user)
	return c.JSON(&user)
}
func DeleteUser(c *fiber.Ctx) error {

}
func UpdateUser(c *fiber.Ctx) error {

}
