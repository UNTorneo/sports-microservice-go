package mode

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mode struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	SportId           primitive.ObjectID `bson:"sportId,omitempty"validate:"required"`
	Name              string             `bson:"name,omitempty"validate:"required"`
	WinningPoints     int                `bson:"winningPoints,omitempty"validate:"required,number"`
	TeamsNumber       int                `bson:"teamsNumber,omitempty"validate:"required,number"`
	PlayersPerTeam    int                `bson:"playersPerTeam,omitempty"validate:"required,number"`
	Description       string             `bson:"description,omitempty"validate:"required"`
	SubstitutePlayers int                `bson:"substitutePlayers,omitempty"validate:"required,number"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateMode(mode Mode) string {
	errors := ""
	err := validate.Struct(mode)
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
