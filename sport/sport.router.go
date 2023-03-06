package sport

import (
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {

	sport := app.Group("/sport")

	sport.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from API!\n")
	})

	sport.Post("/", func(c *fiber.Ctx) error {
		CreateSport(c)
		return c.SendString("jijijijs")
	})

}
