package main

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/jmorenohj/sports/common/config/db"
	"github.com/jmorenohj/sports/sport"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("Hello, World!\n")

	})
	db.Initdb()
	sport.Route(app)
	app.Listen(":3000")
}
