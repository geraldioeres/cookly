package reviews

import "cookly/models/users"

type Review struct {
	ID           int `json:"id"`
	UserID       int `json:"userId"`
	User         users.User
	Rating       int    `json:"rating"`
	RecipeReview string `json:"review"`
}
