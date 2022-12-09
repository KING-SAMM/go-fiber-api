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
	var users []User
	DB.Find(&users)
	return c.JSON(&users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.Find(&user, id)
	return c.JSON(&user)
}

func SaveUser(c *fiber.Ctx) error {
	user := new(User)
	// Get data from request body
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&user)
	return c.JSON(&user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User not available")
	}
	DB.Delete(&user)
	return c.SendString("User deleted successfully!")
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(User)
	DB.First(&user, id)

	// If data does not exist in database
	if user.Email == "" {
		return c.Status(500).SendString("User not available")
	}

	// Get data from request body
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	DB.Save(&user)
	return c.JSON(&user)
}
