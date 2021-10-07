package reviews

import (
	"cookly/mvc/models/recipes"
	"cookly/mvc/models/users"
)

type Review struct {
	ID           int            `json:"id"`
	UserID       int            `json:"userId"`
	User         users.User     `json:"user"`
	RecipeID     int            `json:"recipeId"`
	Recipe       recipes.Recipe `json:"recipe"`
	Rating       float64        `json:"rating"`
	RecipeReview string         `json:"review"`
}
