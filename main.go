package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/KING-SAMM/go-fiber-api/user"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome to studioeternal")
}

func Routers(app *fiber.App) {
	app.Get("/users", user.GetUsers)
	app.Get("/user/:id", user.GetUser)
	app.Post("/user", user.SaveUser)
	app.Delete("/user/:id", user.DeleteUser)
	app.Put("/user/:id", user.UpdateUser)
}

func main() {
	user.InitialMigration()
	app := fiber.New()

	app.Get("/", hello)
	Routers(app)

	app.Listen(":3000")
}