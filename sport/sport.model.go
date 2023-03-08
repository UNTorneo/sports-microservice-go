package sport

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sport struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	Modes          []primitive.ObjectID `bson:"modes,omitempty"validate:"required"`
	Name           string               `bson:"name,omitempty"validate:"required"`
	Description    string               `bson:"description,omitempty"validate:"required"`
	imgs           []string             `bson:"imgs,omitempty"`
	Logo           string               `bson:"logo,omitempty"`
	Recommendation []string             `bson:"recommendation,omitempty"validate:"required"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateSport(sport Sport) string {
	errors := ""
	err := validate.Struct(sport)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = element.Tag + " " + element.FailedField
			return errors
		}
	}
	return errors
}
