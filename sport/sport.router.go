package sport

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jmorenohj/sports/common/config/db"
	env "github.com/jmorenohj/sports/common/config/envs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Route(app *fiber.App) {

	sport := app.Group("/sport")
	sportsCollection := db.Client.Database(env.EnvVariable("CUR_DB")).Collection("sports")

	sport.Get("/:id", func(c *fiber.Ctx) error {

		sportId := c.Params("id")
		objectId, err := primitive.ObjectIDFromHex(sportId)
		query, err := sportsCollection.Find(context.TODO(), bson.D{{"_id", objectId}})

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "El deporte no se encuentra en la base de datos"})
		}

		var sport []bson.M

		if err = query.All(context.TODO(), &sport); err != nil {
			panic(err)
		}

		fmt.Println(sport)

		return c.Status(fiber.StatusOK).JSON(sport[0])
	})

	sport.Post("/", func(c *fiber.Ctx) error {
		sport := new(Sport)

		if err := c.BodyParser(sport); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al guardar el deporte en base de datos"})
		}

		errors := ValidateSport(*sport)
		if errors != "nil" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors})
		}

		results, err := sportsCollection.InsertOne(context.TODO(), sport)
		_ = results
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al guardar el deporte en base de datos"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Deporte añadido exitosamente"})
	})

}
