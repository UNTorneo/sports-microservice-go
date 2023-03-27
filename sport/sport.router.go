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

	sport.Get("/all", func(c *fiber.Ctx) error {
		fmt.Println("Get sports")

		query, err := sportsCollection.Find(context.TODO(), bson.D{})

		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": "Error al cargar los deportes"})
		}

		var sport []bson.M

		if err = query.All(context.TODO(), &sport); err != nil {
			panic(err)
		}

		fmt.Println(sport)

		return c.Status(fiber.StatusOK).JSON(sport)
	})

	sport.Get("/:id", func(c *fiber.Ctx) error {
		fmt.Println("Get sport")
		sportId := c.Params("id")
		objectId, err := primitive.ObjectIDFromHex(sportId)
		query, err := sportsCollection.Find(context.TODO(), bson.D{{"_id", objectId}})

		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": "El deporte no se encuentra en la base de datos"})
		}

		var sport []bson.M

		if err = query.All(context.TODO(), &sport); err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": "Error al recibir la información de la base de datos"})
		}

		fmt.Println(sport)

		return c.Status(fiber.StatusOK).JSON(sport[0])
	})

	sport.Post("/", func(c *fiber.Ctx) error {
		sport := new(Sport)
		fmt.Println("Add sport")
		if err := c.BodyParser(sport); err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": "Error al guardar el deporte en base de datos"})
		}

		errors := ValidateSport(*sport)
		if errors != "" {
			fmt.Println(errors)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": errors})
		}

		results, err := sportsCollection.InsertOne(context.TODO(), sport)
		_ = results
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": "Error al guardar el deporte en base de datos"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Deporte añadido exitosamente"})
	})

	sport.Put("/:id", func(c *fiber.Ctx) error {
		sport := new(Sport)
		sportId := c.Params("id")
		objectId, err := primitive.ObjectIDFromHex(sportId)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": "Error al recibir al actualizar el deporte correspondiente"})
		}

		if err := c.BodyParser(sport); err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": "Error al recibir el deporte en el body"})
		}
		fmt.Println(sport)

		results, err := sportsCollection.UpdateOne(context.TODO(), bson.M{"_id": objectId},
			bson.D{
				{"$set", sport},
			})
		_ = results
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": "Error al guardar el deporte en base de datos"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Deporte actualizado exitosamente"})
	})

	sport.Delete("/:id", func(c *fiber.Ctx) error {

		sportId := c.Params("id")
		objectId, err := primitive.ObjectIDFromHex(sportId)
		query, err := sportsCollection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
		_ = query
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": "El deporte no se encuentra en la base de datos"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Deporte borrado exitosamente"})
	})
}
