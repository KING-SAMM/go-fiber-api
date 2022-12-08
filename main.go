package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/KING-SAMM/go-fiber-api/user"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome to studioeternal")
}

func main() {
	user.initialMigration()
	app := fiber.New()

	app.Get("/", hello)

	app.Listen(":3000")
}