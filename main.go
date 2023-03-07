package main

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/jmorenohj/sports/common/config/db"
	"github.com/jmorenohj/sports/mode"
	"github.com/jmorenohj/sports/sport"
)

func main() {
	app := fiber.New()

	db.Initdb()
	sport.Route(app)
	mode.Route(app)
	app.Listen(":3000")
}
