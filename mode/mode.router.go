package mode

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmorenohj/sports/common/config/db"
	env "github.com/jmorenohj/sports/common/config/envs"
)

func Route(app *fiber.App) {

	mode := app.Group("/mode")
	_ = mode

	modesCollection := db.Client.Database(env.EnvVariable("CUR_DB")).Collection("modes")

	_ = modesCollection

	/*
		sport.Get("/:id", func(c *fiber.Ctx) error {


		})

		sport.Post("/", func(c *fiber.Ctx) error {

		})*/

}
