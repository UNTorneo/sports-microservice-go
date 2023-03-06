package sport

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jmorenohj/sports/common/config/db"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateSport(c *fiber.Ctx) {

	usersCollection := db.Client.Database("testing").Collection("sports")
	_ = usersCollection
	users := []interface{}{
		bson.D{{"fullName", "User 2"}, {"age", 25}},
		bson.D{{"fullName", "User 3"}, {"age", 20}},
		bson.D{{"fullName", "User 4"}, {"age", 28}},
	}
	// insert the bson object slice using InsertMany()
	results, err := usersCollection.InsertMany(context.TODO(), users)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the ids of the newly inserted objects
	fmt.Println(results.InsertedIDs)

}
